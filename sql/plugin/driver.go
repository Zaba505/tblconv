package plugin

import (
	"context"
	"database/sql/driver"
	"errors"
	"io"
	"os/exec"
	"sync"
	"time"

	pb "github.com/Zaba505/tblconv/sql/plugin/proto"

	"github.com/hashicorp/go-plugin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var (
	ErrTransactionAlreadyInProgress = errors.New("plugin: transaction already in progress")
	ErrTransactionAlreadyClosed     = errors.New("plugin: transaction already committed or rolled back")
)

// SQLDriver
type SQLDriver struct {
	prefix     string
	pluginName string
	fullName   string
	envs       []string
	args       []string

	client *plugin.Client
}

// Option
type Option func(*SQLDriver)

// WithPrefix allows users to customize the binary name prefix for plugins.
func WithPrefix(prefix string) Option {
	return func(d *SQLDriver) {
		d.prefix = prefix
	}
}

// WithEnv
func WithEnv(envs ...string) Option {
	return func(d *SQLDriver) {
		d.envs = envs
	}
}

// WithEnv
func WithArgs(args ...string) Option {
	return func(d *SQLDriver) {
		d.args = args
	}
}

// NewDriver
func NewDriver(name string, opts ...Option) *SQLDriver {
	// TODO: sanitize name to prevent shell injection attacks?
	d := &SQLDriver{
		pluginName: name,
	}

	// set options
	for _, opt := range opts {
		opt(d)
	}

	// finalize initialization
	d.fullName = d.prefix + d.pluginName
	cmd := exec.Command(d.fullName, d.args...)
	cmd.Env = d.envs

	d.client = plugin.NewClient(NewClientConfig(cmd))

	return d
}

// Driver returns nil as SQLDriver does not actually implement the driver.Driver interface.
func (d *SQLDriver) Driver() driver.Driver {
	return nil
}

// Close will clean up by waiting for the plugin process to shutdown.
func (d *SQLDriver) Close() error {
	d.client.Kill()
	return nil
}

// Connect
func (d *SQLDriver) Connect(ctx context.Context) (driver.Conn, error) {
	rpcClient, err := d.client.Client()
	if err != nil {
		return nil, err
	}

	// request grpc client for driver plugin
	raw, err := rpcClient.Dispense("driver")
	if err != nil {
		return nil, err
	}

	conn, ok := raw.(driver.Conn)
	if !ok {
		panic("plugin doesn't implement sql/driver.Conn")
	}

	return conn, nil
}

var handshakeConfig = plugin.HandshakeConfig{
	ProtocolVersion:  1,
	MagicCookieKey:   "BASIC_PLUGIN",
	MagicCookieValue: "hello",
}

// NewClientConfig returns a plugin client config.
func NewClientConfig(cmd *exec.Cmd, opts ...grpc.DialOption) *plugin.ClientConfig {
	logger, err := zap.NewDevelopment(zap.IncreaseLevel(zapcore.ErrorLevel))
	if err != nil {
		panic(err)
	}

	return &plugin.ClientConfig{
		HandshakeConfig: handshakeConfig,
		Plugins: map[string]plugin.Plugin{
			"driver": &grpcPlugin{},
		},
		Cmd:              cmd,
		AllowedProtocols: []plugin.Protocol{plugin.ProtocolGRPC},
		Logger:           newLogger(logger.Sugar()),
		AutoMTLS:         true,
		GRPCDialOptions:  opts,
	}
}

// Serve provides a simpler interface for serving a gRPC Driver via
// Hashicorp go-plugin system.
//
func Serve(s pb.DriverServer) {
	plugin.Serve(NewServeConfig(s))
}

// NewServeConfig returns plugin server config.
func NewServeConfig(srvr pb.DriverServer) *plugin.ServeConfig {
	logger, err := zap.NewDevelopment(zap.IncreaseLevel(zapcore.ErrorLevel))
	if err != nil {
		panic(err)
	}

	return &plugin.ServeConfig{
		HandshakeConfig: handshakeConfig,
		Plugins: map[string]plugin.Plugin{
			"driver": &grpcPlugin{Driver: srvr},
		},
		GRPCServer: plugin.DefaultGRPCServer,
		Logger:     newLogger(logger.Sugar()),
	}
}

type grpcPlugin struct {
	plugin.Plugin

	Driver pb.DriverServer
}

func (p *grpcPlugin) GRPCServer(broker *plugin.GRPCBroker, s *grpc.Server) error {
	pb.RegisterDriverServer(s, p.Driver)
	return nil
}

func (p *grpcPlugin) GRPCClient(ctx context.Context, broker *plugin.GRPCBroker, c *grpc.ClientConn) (interface{}, error) {
	return &conn{client: pb.NewDriverClient(c)}, nil
}

// confirm conn implements desired interfaces
var _ interface {
	driver.Conn
} = &conn{}

// TODO: impl Pinger, SessionResetter, and Validator, per database/sql/driver overview.
// TODO: impl ExecerContext, QueryerContext, ConnPrepareContext, ConnBeginTx
type conn struct {
	client pb.DriverClient

	// optional field for signalling that a txn is currently under
	// way on this connection
	mu     sync.Mutex
	txnCtx *pb.TxnContext
}

func (c *conn) Close() error {
	return nil
}

type stmt struct {
	conn *conn

	query string
}

func (c *conn) Prepare(query string) (driver.Stmt, error) {
	return &stmt{conn: c, query: query}, nil
}

func (s *stmt) Close() error {
	return nil
}

func (s *stmt) NumInput() int {
	return -1
}

func (s *stmt) Exec(args []driver.Value) (driver.Result, error) {
	return s.ExecContext(context.Background(), convertValuesToNamedValues(args))
}

func (s *stmt) ExecContext(ctx context.Context, args []driver.NamedValue) (driver.Result, error) {
	return s.do(ctx, args, false)
}

func (s *stmt) Query(args []driver.Value) (driver.Rows, error) {
	return s.QueryContext(context.Background(), convertValuesToNamedValues(args))
}

func (s *stmt) QueryContext(ctx context.Context, args []driver.NamedValue) (driver.Rows, error) {
	return s.do(ctx, args, true)
}

func (s *stmt) do(ctx context.Context, args []driver.NamedValue, returnsRows bool) (*result, error) {
	vals := make([]*pb.NamedValue, 0, len(args))
	for _, arg := range args {
		val := &pb.NamedValue{
			Name:    arg.Name,
			Ordinal: int64(arg.Ordinal),
			Value:   &pb.Value{},
		}
		setValue(val.Value, arg.Value)

		vals = append(vals, val)
	}

	req := &pb.Request{
		StartTs:     uint64(time.Now().UnixNano()),
		Query:       s.query,
		Args:        vals,
		ReturnsRows: returnsRows,
		Txn:         s.conn.txnCtx,
	}

	resp, err := s.conn.client.Query(ctx, req)
	if err != nil {
		return nil, err
	}

	return &result{resp: resp}, nil
}

type tx struct {
	conn *conn
}

func (c *conn) Begin() (driver.Tx, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.txnCtx != nil {
		return nil, ErrTransactionAlreadyInProgress
	}

	c.txnCtx = &pb.TxnContext{}
	return &tx{
		conn: c,
	}, nil
}

func (c *conn) endTx(commit bool) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.txnCtx == nil {
		return ErrTransactionAlreadyClosed
	}
	_, err := c.client.CommitOrRollback(context.Background(), c.txnCtx)
	c.txnCtx = nil
	return err
}

func (t *tx) Commit() error {
	return t.conn.endTx(true)
}

func (t *tx) Rollback() error {
	return t.conn.endTx(false)
}

func convertValuesToNamedValues(vals []driver.Value) []driver.NamedValue {
	namedVals := make([]driver.NamedValue, 0, len(vals))
	for i, val := range vals {
		namedVals = append(namedVals, driver.NamedValue{
			Ordinal: i + 1,
			Value:   val,
		})
	}
	return namedVals
}

type result struct {
	resp *pb.Response

	rowIdx int
}

func (r *result) LastInsertId() (int64, error) {
	return r.resp.LastInsertId, nil
}

func (r *result) RowsAffected() (int64, error) {
	return r.resp.RowsAffected, nil
}

func (r *result) Close() error {
	return nil
}

func (r *result) Columns() []string {
	return r.resp.Columns
}

func (r *result) Next(dest []driver.Value) error {
	if r.rowIdx >= len(r.resp.Rows) {
		return io.EOF
	}

	row := r.resp.Rows[r.rowIdx]
	for i, col := range row.Columns {
		dest[i] = getValue(col.Value)
	}
	r.rowIdx += 1
	return nil
}

// set value based on one of 7 builtin types of the sql package.
func setValue(val *pb.Value, v any) {
	switch x := v.(type) {
	case nil:
		val.Value = &pb.Value_Null{
			Null: true,
		}
	case int64:
		val.Value = &pb.Value_Int64{
			Int64: x,
		}
	case float64:
		val.Value = &pb.Value_Float64{
			Float64: x,
		}
	case bool:
		val.Value = &pb.Value_Bool{
			Bool: x,
		}
	case []byte:
		val.Value = &pb.Value_Bytes{
			Bytes: x,
		}
	case string:
		val.Value = &pb.Value_String_{
			String_: x,
		}
	case time.Time:
		val.Value = &pb.Value_Time{
			Time: timestamppb.New(x),
		}
	default:
		panic("unsupported value type")
	}
}

// return one of the 7 builtin types of the sql package
func getValue(val *pb.Value) any {
	switch x := val.Value.(type) {
	case *pb.Value_Null:
		return nil
	case *pb.Value_Int64:
		return x.Int64
	case *pb.Value_Float64:
		return x.Float64
	case *pb.Value_Bool:
		return x.Bool
	case *pb.Value_Bytes:
		return x.Bytes
	case *pb.Value_String_:
		return x.String_
	case *pb.Value_Time:
		return x.Time.AsTime()
	default:
		panic("unrecognized value type")
	}
}
