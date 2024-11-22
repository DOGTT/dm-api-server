package server

import (
	"context"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/DOGTT/dm-api-server/internal/conf"
	"github.com/DOGTT/dm-api-server/internal/service"
	log "github.com/uptrace/opentelemetry-go-extra/otelzap"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

// Server type.
type Server struct {
	c          *conf.Server
	httpServer *http.Server
	grpcServer *grpc.Server
	svc        *service.Service
}

// New Server instance.
func New(c *conf.Server, svc *service.Service) (*Server, error) {
	s := &Server{
		c:   c,
		svc: svc,
	}
	if s.c.GRPC.Enable {
		s.grpcServer = NewGRPCServer(c, s.svc)
	}
	if s.c.HTTP.Enable {
		s.httpServer = NewGinServer(c, s.svc)
	}
	return s, nil
}

// Start the server, include grpc server and http server.
func (s *Server) Start() {
	defer s.Stop()
	ch := make(chan error, 2)

	// start grpc server
	if s.grpcServer != nil {
		go s.startGrpcServer(ch)
	}
	// start http server
	if s.httpServer != nil {
		go s.startHTTPServer(ch)
	}

	sigs := make(chan os.Signal, 3)
	signal.Notify(sigs, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	select {
	case sig := <-sigs:
		log.L().Info("receive signal, server exited", zap.Any("signal", sig))
	case err := <-ch:
		log.L().Panic("run server failed", zap.Error(err))
	}
}

// Stop the server, include grpc and http server.
func (s *Server) Stop() {
	if s.grpcServer != nil {
		s.grpcServer.GracefulStop()
	}
	if s.httpServer != nil {
		if err := s.httpServer.Shutdown(context.Background()); err != nil {
			log.L().Warn("http sever stop failed", zap.Error(err))
		}
	}
	_ = zap.L().Sync()
}

func (s *Server) startGrpcServer(ch chan<- error) {
	lis, err := net.Listen("tcp", s.c.GRPC.Addr)
	if err != nil {
		log.L().Error("grpc listen failed", zap.Error(err))
		ch <- err
		return
	}
	log.L().Info("GRPC Listening", zap.Any("addr", s.c.GRPC.Addr))
	ch <- s.grpcServer.Serve(lis)
}

func (s *Server) startHTTPServer(ch chan<- error) {
	s.httpServer.Addr = s.c.HTTP.Addr
	log.L().Info("HTTP Listening", zap.Any("addr", s.c.HTTP.Addr))
	ch <- s.httpServer.ListenAndServe()
}
