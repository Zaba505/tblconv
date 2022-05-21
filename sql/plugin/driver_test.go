package driver

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

func TestDriver(t *testing.T) {
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
		if !assert.Nil(subT, err) || !assert.Equal(subT, 1, lastInsertId) {
			return
		}

		rowsAffected, err := result.RowsAffected()
		if !assert.Nil(subT, err) || !assert.Equal(subT, 1, rowsAffected) {
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
		defer rows.Close()

		cols, err := rows.Columns()
		if !assert.Nil(subT, err) || !assert.Equal(subT, 1, len(cols)) {
			return
		}

		colTypes, err := rows.ColumnTypes()
		if !assert.Nil(subT, err) || !assert.Equal(subT, 1, len(colTypes)) {
			return
		}
		if !assert.Equal(subT, "VARCHAR", colTypes[0].DatabaseTypeName()) {
			return
		}

		vals := make([]string, 0, 10)
		for rows.Next() {
			var val string
			if err := rows.Scan(&val); !assert.Nil(subT, err) {
				return
			}
			vals = append(vals, val)
		}
		if err := rows.Err(); !assert.Nil(subT, err) {
			return
		}

		if !assert.Equal(subT, 10, len(vals)) {
			return
		}
	})

	t.Run("should be able to execute a prepared statement", func(subT *testing.T) {
		subT.Fail()
	})

	t.Run("should be able to query with a prepared statement", func(subT *testing.T) {
		subT.Fail()
	})

	t.Run("should be able to start a transaction", func(subT *testing.T) {
		subT.Fail()
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
				"driver": &testGrpcPlugin{},
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
				"driver": &testGrpcPlugin{},
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
}

func (p *testGrpcPlugin) GRPCServer(broker *plugin.GRPCBroker, s *grpc.Server) error {
	pb.RegisterDriverServer(s, &testDriverServer{})
	return nil
}

func (p *testGrpcPlugin) GRPCClient(ctx context.Context, broker *plugin.GRPCBroker, c *grpc.ClientConn) (interface{}, error) {
	return nil, nil
}

type testDriverServer struct {
	pb.UnimplementedDriverServer
}

func (*testDriverServer) Login(ctx context.Context, req *pb.LoginRequest) (*pb.Response, error) {
	return new(pb.Response), nil
}

func (*testDriverServer) Query(ctx context.Context, req *pb.Request) (*pb.Response, error) {
	return new(pb.Response), nil
}

func (*testDriverServer) CommitOrAbort(ctx context.Context, req *pb.TxnContext) (*pb.TxnContext, error) {
	return new(pb.TxnContext), nil
}
