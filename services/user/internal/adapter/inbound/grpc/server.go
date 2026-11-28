package grpc

import (
	"context"
	"errors"
	"log/slog"
	"net"

	userv1 "github.com/LarsSeverson/charter/gen/charter/user/v1"
	"github.com/LarsSeverson/charter/services/user/internal/adapter/inbound/grpc/interceptor"
	"google.golang.org/grpc"
)

type Server struct {
	grpcServer *grpc.Server
	address    string
	listener   net.Listener
	fail       func(error)
}

// New builds the server and registers the handlers. It performs no I/O;
// the address is bound in Start.
//
// fail is called if the server stops serving unexpectedly after Start
// has returned.
func New(
	config Config,
	userService userv1.UserServiceServer,
	logger *slog.Logger,
	fail func(error),
) *Server {
	unaryInterceptors := serverUnaryInterceptors(logger, config)
	options := serverOptions(unaryInterceptors, config)
	grpcServer := grpc.NewServer(options...)

	userv1.RegisterUserServiceServer(
		grpcServer,
		userService,
	)

	server := &Server{
		grpcServer: grpcServer,
		address:    config.Address,
		fail:       fail,
	}
	return server
}

// Lifecycle

func (s *Server) Name() string {
	return "grpc"
}

func (s *Server) Start(ctx context.Context) error {
	listenConfig := net.ListenConfig{}

	listener, err := listenConfig.Listen(ctx, "tcp", s.address)
	if err != nil {
		return errors.Join(ErrListen, err)
	}

	s.listener = listener

	go s.serve()

	return nil
}

// Stop drains in-flight requests. If ctx expires first, remaining
// connections are closed forcefully.
func (s *Server) Stop(ctx context.Context) error {
	stopped := make(chan struct{})

	go func() {
		s.grpcServer.GracefulStop()
		close(stopped)
	}()

	select {
	case <-stopped:
		return nil
	case <-ctx.Done():
		s.grpcServer.Stop()
		<-stopped

		err := errors.Join(ErrForceStop, ctx.Err())
		return err
	}
}

func (s *Server) serve() {
	err := s.grpcServer.Serve(s.listener)
	if err == nil {
		return
	}

	serveErr := errors.Join(ErrServe, err)
	s.fail(serveErr)
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
	n := len(config.Options) + 1
	options := make([]grpc.ServerOption, 0, n)
	options = append(options, grpc.ChainUnaryInterceptor(unaryInterceptors...))
	options = append(options, config.Options...)
	return options
}
