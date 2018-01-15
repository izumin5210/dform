package repo

import (
	"context"

	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	"go.uber.org/zap"
	"google.golang.org/grpc"

	"github.com/izumin5210/dform/util/log"
)

// GrpcPool pools gRPC client connections
type GrpcPool interface {
	Get() (*grpc.ClientConn, error)
}

// NewGrpcPool creates GrpcPool instance
func NewGrpcPool(url string) GrpcPool {
	return &grpcPool{
		url: url,
	}
}

type grpcPool struct {
	url string
}

func (p *grpcPool) Get() (*grpc.ClientConn, error) {
	return grpc.Dial(p.url, dialOptions()...)
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
