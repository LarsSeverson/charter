package server

import (
	"context"
	"errors"
	"log/slog"
	"net"
	"time"

	userv1 "github.com/LarsSeverson/charter/gen/charter/user/v1"
	"github.com/LarsSeverson/charter/services/user/internal/transport/grpc/interceptor"
	"google.golang.org/grpc"
)

type Server struct {
	grpcServer      *grpc.Server
	listener        net.Listener
	shutdownTimeout time.Duration
}

func New(
	listener net.Listener,
	userService userv1.UserServiceServer,
	logger *slog.Logger,
	config Config,
) *Server {
	unaryInterceptors := serverUnaryInterceptors(logger, config)
	options := serverOptions(unaryInterceptors, config)
	grpcServer := grpc.NewServer(options...)

	userv1.RegisterUserServiceServer(
		grpcServer,
		userService,
	)

	return &Server{
		grpcServer:      grpcServer,
		listener:        listener,
		shutdownTimeout: config.ShutdownTimeout,
	}
}

func (s *Server) Run(ctx context.Context) error {
	serveErrors := make(chan error, 1)

	go func() {
		serveErrors <- s.grpcServer.Serve(s.listener)
	}()

	select {
	case err := <-serveErrors:
		if err != nil {
			return errors.Join(ErrServe, err)
		}

		return nil

	case <-ctx.Done():
		s.shutdown()

		if err := <-serveErrors; err != nil {
			return errors.Join(ErrServe, err)
		}

		return nil
	}
}

func (s *Server) shutdown() {
	stopped := make(chan struct{})

	go func() {
		s.grpcServer.GracefulStop()
		close(stopped)
	}()

	timer := time.NewTimer(s.shutdownTimeout)
	defer timer.Stop()

	select {
	case <-stopped:
		return
	case <-timer.C:
		s.grpcServer.Stop()
		<-stopped
	}
}

func serverUnaryInterceptors(
	logger *slog.Logger,
	config Config,
) []grpc.UnaryServerInterceptor {
	unaryInterceptors := []grpc.UnaryServerInterceptor{
		interceptor.RequestID(),
		interceptor.Recovery(logger),
		interceptor.Logging(logger),
	}
	unaryInterceptors = append(unaryInterceptors, config.UnaryInterceptors...)
	return unaryInterceptors
}

func serverOptions(
	unaryInterceptors []grpc.UnaryServerInterceptor,
	config Config,
) []grpc.ServerOption {
	options := make([]grpc.ServerOption, 0, len(config.Options)+1)
	options = append(options, grpc.ChainUnaryInterceptor(unaryInterceptors...))
	options = append(options, config.Options...)
	return options
}
