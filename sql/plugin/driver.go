package driver

import (
	"context"
	"database/sql/driver"
	"os/exec"

	pb "github.com/Zaba505/tblconv/sql/plugin/proto"

	"github.com/hashicorp/go-plugin"
	"google.golang.org/grpc"
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

	d.client = plugin.NewClient(&plugin.ClientConfig{
		HandshakeConfig: plugin.HandshakeConfig{
			ProtocolVersion:  1,
			MagicCookieKey:   "BASIC_PLUGIN",
			MagicCookieValue: "hello",
		},
		Plugins: map[string]plugin.Plugin{
			"driver": &driverPlugin{},
		},
		Cmd:              cmd,
		AllowedProtocols: []plugin.Protocol{plugin.ProtocolGRPC},
		AutoMTLS:         true,
	})

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

type driverPlugin struct {
	plugin.Plugin
}

func (p *driverPlugin) GRPCServer(broker *plugin.GRPCBroker, s *grpc.Server) error {
	return nil
}

func (p *driverPlugin) GRPCClient(ctx context.Context, broker *plugin.GRPCBroker, c *grpc.ClientConn) (interface{}, error) {
	return &conn{client: pb.NewDriverClient(c)}, nil
}

// confirm conn implements driver.Conn to avoid panicing in SQLDriver.Connect
var _ driver.Conn = &conn{}

type conn struct {
	client pb.DriverClient
}

func (c *conn) Prepare(query string) (driver.Stmt, error) {
	return nil, nil
}

func (c *conn) Close() error {
	return nil
}

func (c *conn) Begin() (driver.Tx, error) {
	return nil, nil
}
