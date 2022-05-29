package cmd

import (
	"context"
	"database/sql"
	"time"

	pb "github.com/Zaba505/tblconv/sql/plugin/proto"

	"github.com/hashicorp/go-plugin"
	_ "github.com/mattn/go-sqlite3"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func Execute(ctx context.Context) {
	db, err := sql.Open("sqlite3", "./foo.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: plugin.HandshakeConfig{
			ProtocolVersion:  1,
			MagicCookieKey:   "BASIC_PLUGIN",
			MagicCookieValue: "hello",
		},
		Plugins: map[string]plugin.Plugin{
			"driver": &sqlitePlugin{
				db: db,
			},
		},

		// A non-nil value here enables gRPC serving for this plugin...
		GRPCServer: plugin.DefaultGRPCServer,
	})
	return
}

type sqlitePlugin struct {
	plugin.Plugin
	pb.UnimplementedDriverServer

	db *sql.DB
}

func (p *sqlitePlugin) GRPCServer(broker *plugin.GRPCBroker, s *grpc.Server) error {
	pb.RegisterDriverServer(s, p)
	return nil
}

func (p *sqlitePlugin) GRPCClient(ctx context.Context, broker *plugin.GRPCBroker, c *grpc.ClientConn) (interface{}, error) {
	return nil, nil
}

func (p *sqlitePlugin) Query(ctx context.Context, req *pb.Request) (*pb.Response, error) {
	if req.ReturnsRows {
		return p.query(ctx, req)
	}
	return p.exec(ctx, req)
}

func (p *sqlitePlugin) query(ctx context.Context, req *pb.Request) (*pb.Response, error) {
	rows, err := p.db.QueryContext(ctx, req.Query, getRawValues(req.Args)...)
	if err != nil {
		return nil, err
	}

	cols, err := rows.Columns()
	if err != nil {
		return nil, err
	}
	resp := &pb.Response{
		Columns: cols,
	}
	for rows.Next() {
		vals := make([]any, len(cols))
		results := make([]any, len(cols))
		for i := range vals {
			results[i] = &vals[i]
		}
		if err := rows.Scan(results...); err != nil {
			return nil, err
		}
		columns := make([]*pb.Column, len(cols))
		for i := range cols {
			columns[i] = &pb.Column{
				Name:  cols[i],
				Value: convertRawToValue(vals[i]),
			}
		}
		resp.Rows = append(resp.Rows, &pb.Row{
			Columns: columns,
		})
	}
	return resp, nil
}

func (p *sqlitePlugin) exec(ctx context.Context, req *pb.Request) (*pb.Response, error) {
	res, err := p.db.ExecContext(ctx, req.Query, getRawValues(req.Args)...)
	if err != nil {
		return nil, err
	}

	lastId, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return nil, err
	}
	resp := &pb.Response{
		LastInsertId: lastId,
		RowsAffected: rowsAffected,
	}
	return resp, nil
}

func (p *sqlitePlugin) CommitOrRollback(ctx context.Context, req *pb.TxnContext) (*pb.TxnContext, error) {
	// TODO: actually commit txn
	return req, nil
}

func getRawValues(args []*pb.NamedValue) []any {
	rawVals := make([]any, 0, len(args))
	for _, arg := range args {
		rawVals = append(rawVals, getRawValue(arg.Value))
	}
	return rawVals
}

func getRawValue(val *pb.Value) any {
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
		return nil
	}
}

func convertRawToValue(v any) *pb.Value {
	switch x := v.(type) {
	case nil:
		return &pb.Value{
			Value: &pb.Value_Null{Null: true},
		}
	case int64:
		return &pb.Value{
			Value: &pb.Value_Int64{Int64: x},
		}
	case float64:
		return &pb.Value{
			Value: &pb.Value_Float64{Float64: x},
		}
	case bool:
		return &pb.Value{
			Value: &pb.Value_Bool{Bool: x},
		}
	case []byte:
		return &pb.Value{
			Value: &pb.Value_Bytes{Bytes: x},
		}
	case string:
		return &pb.Value{
			Value: &pb.Value_String_{String_: x},
		}
	case time.Time:
		return &pb.Value{
			Value: &pb.Value_Time{Time: timestamppb.New(x)},
		}
	default:
		return nil
	}
}
