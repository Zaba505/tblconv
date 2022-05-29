package plugin

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"testing"
	"time"

	pb "github.com/Zaba505/tblconv/sql/plugin/proto"

	"github.com/hashicorp/go-plugin"
	flag "github.com/spf13/pflag"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
)

func assertSuccessfulClose(t assert.TestingT, f func() error) bool {
	return assert.Nil(t, f())
}

func TestSQLDriver(t *testing.T) {

	//
	// General DB operations
	//

	t.Run("should fail to connect if plugin binary can not be found", func(subT *testing.T) {
		d := NewDriver("nonexistent-plugin")
		_, err := d.Connect(context.Background())
		if !assert.Error(subT, err) {
			return
		}

		e := errors.Unwrap(err)
		if !assert.Equal(subT, exec.ErrNotFound, e) {
			return
		}
	})

	t.Run("should successfully be able to ping and close db", func(subT *testing.T) {
		args := getHelperPluginCLI("pingable")
		d := NewDriver(args[0], WithArgs(args[1:]...), WithEnv("GO_WANT_HELPER_PROCESS=1"))
		db := sql.OpenDB(d)

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		err := db.PingContext(ctx)
		if !assert.Nil(subT, err) {
			return
		}

		err = db.Close()
		if !assert.Nil(subT, err) {
			return
		}
	})

	//
	// Non-transaction based operations
	//

	t.Run("should be able to execute a query without returning any rows", func(subT *testing.T) {
		args := getHelperPluginCLI("execute", "--LastInsertId=1", "--RowsAffected=1")
		d := NewDriver(args[0], WithArgs(args[1:]...), WithEnv("GO_WANT_HELPER_PROCESS=1"))
		db := sql.OpenDB(d)

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		result, err := db.ExecContext(ctx, "INSERT INTO t (hello) VALUES world")
		if !assert.Nil(subT, err) {
			return
		}

		lastInsertId, err := result.LastInsertId()
		if !assert.Nil(subT, err) || !assert.Equal(subT, int64(1), lastInsertId) {
			return
		}

		rowsAffected, err := result.RowsAffected()
		if !assert.Nil(subT, err) || !assert.Equal(subT, int64(1), rowsAffected) {
			return
		}
	})

	t.Run("should be able to execute a query with return rows", func(subT *testing.T) {
		args := getHelperPluginCLI("query", "--Columns=HELLO", "--ColumnTypes=VARCHAR", "--TotalRows=10")
		d := NewDriver(args[0], WithArgs(args[1:]...), WithEnv("GO_WANT_HELPER_PROCESS=1"))
		db := sql.OpenDB(d)

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		rows, err := db.QueryContext(ctx, "SELECT * FROM t")
		if !assert.Nil(subT, err) {
			return
		}
		defer assertSuccessfulClose(subT, rows.Close)

		cols, err := rows.Columns()
		if !assert.Nil(subT, err) || !assert.Equal(subT, 1, len(cols)) {
			return
		}

		colTypes, err := rows.ColumnTypes()
		if !assert.Nil(subT, err) || !assert.Equal(subT, 1, len(colTypes)) {
			return
		}

		vals := make([]string, 0, 10)
		for rows.Next() {
			var val sql.NullString
			if err := rows.Scan(&val); !assert.Nil(subT, err) {
				return
			}
			vals = append(vals, val.String)
		}
		if err := rows.Err(); !assert.Nil(subT, err) {
			return
		}

		if !assert.Equal(subT, 10, len(vals)) {
			return
		}
	})

	t.Run("should be able to execute a prepared statement", func(subT *testing.T) {
		args := getHelperPluginCLI("execute", "--LastInsertId=1", "--RowsAffected=1")
		d := NewDriver(args[0], WithArgs(args[1:]...), WithEnv("GO_WANT_HELPER_PROCESS=1"))
		db := sql.OpenDB(d)

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		stmt, err := db.PrepareContext(ctx, "INSERT INTO t (hello) VALUES ?")
		if !assert.Nil(subT, err) {
			return
		}
		defer assertSuccessfulClose(subT, stmt.Close)

		result, err := stmt.ExecContext(ctx, "world")
		if !assert.Nil(subT, err) {
			return
		}

		lastInsertId, err := result.LastInsertId()
		if !assert.Nil(subT, err) || !assert.Equal(subT, int64(1), lastInsertId) {
			return
		}

		rowsAffected, err := result.RowsAffected()
		if !assert.Nil(subT, err) || !assert.Equal(subT, int64(1), rowsAffected) {
			return
		}
	})

	t.Run("should be able to query with a prepared statement", func(subT *testing.T) {
		args := getHelperPluginCLI("query", "--Columns=HELLO", "--ColumnTypes=VARCHAR", "--TotalRows=10")
		d := NewDriver(args[0], WithArgs(args[1:]...), WithEnv("GO_WANT_HELPER_PROCESS=1"))
		db := sql.OpenDB(d)

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		stmt, err := db.PrepareContext(ctx, "SELECT * FROM t WHERE HELLO = ?")
		if !assert.Nil(subT, err) {
			return
		}
		defer assertSuccessfulClose(subT, stmt.Close)

		rows, err := stmt.QueryContext(ctx, "world")
		if !assert.Nil(subT, err) {
			return
		}
		defer assertSuccessfulClose(subT, rows.Close)

		cols, err := rows.Columns()
		if !assert.Nil(subT, err) || !assert.Equal(subT, 1, len(cols)) {
			return
		}

		colTypes, err := rows.ColumnTypes()
		if !assert.Nil(subT, err) || !assert.Equal(subT, 1, len(colTypes)) {
			return
		}

		vals := make([]string, 0, 10)
		for rows.Next() {
			var val sql.NullString
			if err := rows.Scan(&val); !assert.Nil(subT, err) {
				return
			}
			vals = append(vals, val.String)
		}
		if err := rows.Err(); !assert.Nil(subT, err) {
			return
		}

		if !assert.Equal(subT, 10, len(vals)) {
			return
		}
	})
}

func getHelperPluginCLI(s ...string) []string {
	cs := []string{"-test.run=TestHelperProcess", "--"}
	cs = append(cs, s...)
	return append(os.Args[:1], cs...)
}

// TestHelperProcess isn't a real test. It's used as a helper process
// for TestDriver.
//
func TestHelperProcess(t *testing.T) {
	if os.Getenv("GO_WANT_HELPER_PROCESS") != "1" {
		return
	}
	defer os.Exit(0)

	args := os.Args
	for len(args) > 0 {
		if args[0] == "--" {
			args = args[1:]
			break
		}
		args = args[1:]
	}
	if len(args) == 0 {
		fmt.Fprintf(os.Stderr, "No command\n")
		os.Exit(2)
	}

	cmd, args := args[0], args[1:]
	flags := flag.NewFlagSet(cmd, flag.ExitOnError)
	switch cmd {
	case "pingable":
		plugin.Serve(&plugin.ServeConfig{
			HandshakeConfig: plugin.HandshakeConfig{
				ProtocolVersion:  1,
				MagicCookieKey:   "BASIC_PLUGIN",
				MagicCookieValue: "hello",
			},
			Plugins: map[string]plugin.Plugin{
				"driver": &testGrpcPlugin{},
			},

			// A non-nil value here enables gRPC serving for this plugin...
			GRPCServer: plugin.DefaultGRPCServer,
		})
	case "execute":
		var executeFlags struct {
			LastInsertId int64
			RowsAffected int64
		}
		flags.Int64Var(&executeFlags.LastInsertId, "LastInsertId", 0, "")
		flags.Int64Var(&executeFlags.RowsAffected, "RowsAffected", 0, "")
		err := flags.Parse(args)
		if err != nil {
			panic(err)
		}

		plugin.Serve(&plugin.ServeConfig{
			HandshakeConfig: plugin.HandshakeConfig{
				ProtocolVersion:  1,
				MagicCookieKey:   "BASIC_PLUGIN",
				MagicCookieValue: "hello",
			},
			Plugins: map[string]plugin.Plugin{
				"driver": &testGrpcPlugin{
					LastInsertId: executeFlags.LastInsertId,
					RowsAffected: executeFlags.RowsAffected,
				},
			},

			// A non-nil value here enables gRPC serving for this plugin...
			GRPCServer: plugin.DefaultGRPCServer,
		})
	case "query":
		var queryFlags struct {
			Columns     []string
			ColumnTypes []string
			TotalRows   int
		}
		flags.StringSliceVar(&queryFlags.Columns, "Columns", nil, "")
		flags.StringSliceVar(&queryFlags.ColumnTypes, "ColumnTypes", nil, "")
		flags.IntVar(&queryFlags.TotalRows, "TotalRows", 0, "")
		err := flags.Parse(args)
		if err != nil {
			panic(err)
		}

		plugin.Serve(&plugin.ServeConfig{
			HandshakeConfig: plugin.HandshakeConfig{
				ProtocolVersion:  1,
				MagicCookieKey:   "BASIC_PLUGIN",
				MagicCookieValue: "hello",
			},
			Plugins: map[string]plugin.Plugin{
				"driver": &testGrpcPlugin{
					Columns:     queryFlags.Columns,
					ColumnTypes: queryFlags.ColumnTypes,
					TotalRows:   queryFlags.TotalRows,
				},
			},

			// A non-nil value here enables gRPC serving for this plugin...
			GRPCServer: plugin.DefaultGRPCServer,
		})
	default:
		// TODO: fail here
		return
	}
}

type testGrpcPlugin struct {
	plugin.Plugin
	pb.UnimplementedDriverServer

	LastInsertId int64
	RowsAffected int64

	Columns     []string
	ColumnTypes []string
	TotalRows   int
}

func (p *testGrpcPlugin) GRPCServer(broker *plugin.GRPCBroker, s *grpc.Server) error {
	pb.RegisterDriverServer(s, p)
	return nil
}

func (p *testGrpcPlugin) GRPCClient(ctx context.Context, broker *plugin.GRPCBroker, c *grpc.ClientConn) (interface{}, error) {
	return nil, nil
}

func (p *testGrpcPlugin) Query(ctx context.Context, req *pb.Request) (*pb.Response, error) {
	resp := &pb.Response{
		LastInsertId: p.LastInsertId,
		RowsAffected: p.RowsAffected,
		Columns:      p.Columns,
	}
	for i := 0; i < p.TotalRows; i++ {
		resp.Rows = append(resp.Rows, newRow(p.Columns))
	}
	return resp, nil
}

func (p *testGrpcPlugin) CommitOrRollback(ctx context.Context, req *pb.TxnContext) (*pb.TxnContext, error) {
	resp := &pb.TxnContext{
		StartTs:   req.StartTs,
		CommitTs:  time.Now().UnixNano(),
		Committed: req.Committed,
		Aborted:   req.Aborted,
	}
	return resp, nil
}

func newRow(columnNames []string) *pb.Row {
	cols := make([]*pb.Column, 0, len(columnNames))
	for _, name := range columnNames {
		cols = append(cols, &pb.Column{
			Name: name,
			Value: &pb.Value{
				Value: new(pb.Value_Null),
			},
		})
	}
	return &pb.Row{
		Columns: cols,
	}
}
