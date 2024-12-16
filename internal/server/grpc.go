package server

import (
	"context"
	"time"

	api "github.com/DOGTT/dm-api-server/api/base"
	"github.com/DOGTT/dm-api-server/internal/conf"
	"github.com/DOGTT/dm-api-server/internal/service"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
)

type validator interface {
	Validate() error
}

// 创建验证拦截器
func validationInterceptor() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		// 验证请求
		if v, ok := req.(validator); ok {
			if err := v.Validate(); err != nil {
				return nil, err
			}
		}
		// 调用下一个处理程序
		return handler(ctx, req)
	}
}

func NewGRPCServer(c *conf.Server, svc *service.Service) *grpc.Server {
	var opts = []grpc.ServerOption{
		grpc.ConnectionTimeout(time.Second * 30),
	}
	unaryInterceptors := []grpc.UnaryServerInterceptor{validationInterceptor()}
	streamInterceptors := []grpc.StreamServerInterceptor{}
	if c.GRPC.EnableTrace {
		opts = append(opts, grpc.StatsHandler(otelgrpc.NewServerHandler()))
	}
	if c.GRPC.EnableMetric {
		unaryInterceptors = append(unaryInterceptors, grpc_prometheus.UnaryServerInterceptor)
		streamInterceptors = append(streamInterceptors, grpc_prometheus.StreamServerInterceptor)
	}
	opts = append(opts, grpc.ChainUnaryInterceptor(unaryInterceptors...), grpc.ChainStreamInterceptor(streamInterceptors...))

	srv := grpc.NewServer(opts...)

	api.RegisterBaseServiceServer(srv, svc)
	return srv
}
