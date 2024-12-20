package server

import (
	"time"

	api "github.com/DOGTT/dm-api-server/api/grpc"
	"github.com/DOGTT/dm-api-server/internal/conf"
	"github.com/DOGTT/dm-api-server/internal/service"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
)

func NewGRPCServer(c *conf.Server, svc *service.Service) *grpc.Server {
	var opts = []grpc.ServerOption{
		grpc.ConnectionTimeout(time.Second * 30),
	}
	unaryInterceptors := []grpc.UnaryServerInterceptor{}
	streamInterceptors := []grpc.StreamServerInterceptor{}
	if c.GRPC.EnableTrace {
		opts = append(opts, grpc.StatsHandler(otelgrpc.NewServerHandler()))
		// unaryInterceptors = append(unaryInterceptors, otelgrpc.UnaryServerInterceptor())
		// streamInterceptors = append(streamInterceptors, otelgrpc.StreamServerInterceptor())
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
