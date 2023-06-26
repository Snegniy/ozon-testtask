package grpchandlers

import (
	"context"
	"errors"
	"github.com/Snegniy/ozon-testtask/internal/service"
	pb "github.com/Snegniy/ozon-testtask/pkg/api"
	"github.com/Snegniy/ozon-testtask/pkg/logger"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net"
	"os"
	"os/signal"
	"time"
)

type Server struct {
	pb.LinkShortenerServer
	services *service.Service
}

func NewGrpcServer(service *service.Service) *Server {
	logger.Debug("grpc server create")
	return &Server{
		services: service,
	}
}

func (s *Server) StartServer(host string) {
	logger.Debug("Start gRPC app server")

	go func() {
		logger.Info("Server gRPC started", zap.String("host:port", host))
		lis, err := net.Listen("tcp", host)
		if err != nil {
			logger.Fatal("listen:", zap.Error(err))
		}

		grpcServer := grpc.NewServer()

		pb.RegisterLinkShortenerServer(grpcServer, s)
		if err := grpcServer.Serve(lis); err != nil {
			logger.Fatal("failed gRPC server:", zap.Error(err))
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	logger.Info("Shutdown gRPC Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	go func() {
		<-ctx.Done()
		if errors.Is(ctx.Err(), context.DeadlineExceeded) {
			logger.Fatal("Server gRPC Shutdown timeout exit")
		}
	}()
	logger.Info("Server gRPC exiting")

}
