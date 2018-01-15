package component

import (
	"context"

	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	"go.uber.org/zap"
	"google.golang.org/grpc"

	"github.com/izumin5210/dform/app/system"
	"github.com/izumin5210/dform/domain/schema"
	"github.com/izumin5210/dform/infra/repo"
	"github.com/izumin5210/dform/util/log"
)

// Dgraph containes dependencies for accessing to Dgraph.
type Dgraph interface {
	DgraphSchemaRepository() (schema.Repository, error)
}

func newDgraph(config *system.Config) Dgraph {
	return &dgraph{
		config: config,
	}
}

type dgraph struct {
	config     *system.Config
	conn       *grpc.ClientConn
	schemaRepo schema.Repository
}

func (d *dgraph) DgraphSchemaRepository() (schema.Repository, error) {
	if d.schemaRepo != nil {
		return d.schemaRepo, nil
	}
	conn, err := d.getConn()
	if err != nil {
		return nil, err
	}
	d.schemaRepo = repo.NewDgraphSchemaRepository(conn)
	return d.schemaRepo, nil
}

func (d *dgraph) getConn() (*grpc.ClientConn, error) {
	if d.conn != nil {
		return d.conn, nil
	}
	conn, err := grpc.Dial(d.config.GetDgraphURL(), dialOptions()...)
	if err != nil {
		return nil, err
	}
	d.conn = conn
	return d.conn, nil
}

func dialOptions() []grpc.DialOption {
	return []grpc.DialOption{
		grpc.WithInsecure(),
		unaryClientInterceptor(),
		streamClientInterceptor(),
	}
}

func unaryClientInterceptor() grpc.DialOption {
	return grpc.WithUnaryInterceptor(
		grpc_middleware.ChainUnaryClient(
			grpc_zap.UnaryClientInterceptor(clientInterceptorLogger()),
			grpc_zap.PayloadUnaryClientInterceptor(payloadInterceptorLogger(), defaultDecider),
		),
	)
}

func streamClientInterceptor() grpc.DialOption {
	return grpc.WithStreamInterceptor(
		grpc_middleware.ChainStreamClient(
			grpc_zap.StreamClientInterceptor(clientInterceptorLogger()),
			grpc_zap.PayloadStreamClientInterceptor(payloadInterceptorLogger(), defaultDecider),
		),
	)
}

var defaultDecider = func(context.Context, string) bool { return true }

func clientInterceptorLogger() *zap.Logger {
	return log.Logger().WithOptions(zap.AddCallerSkip(2))
}

func payloadInterceptorLogger() *zap.Logger {
	return log.Logger().WithOptions(zap.AddCallerSkip(4))
}
