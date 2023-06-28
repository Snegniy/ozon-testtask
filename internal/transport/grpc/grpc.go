package grpchandlers

import (
	"context"
	"errors"
	"github.com/Snegniy/ozon-testtask/internal/apperror"
	pb "github.com/Snegniy/ozon-testtask/pkg/api"
	"github.com/Snegniy/ozon-testtask/pkg/logger"
	"go.uber.org/zap"
)

func (s *Server) GetShortLink(ctx context.Context, request *pb.GetShortLinkRequest) (*pb.GetShortLinkResponse, error) {
	logger.Info("GetShortLink gRPC request", zap.String("url", request.GetUrl()))
	url, err := s.services.GetShortLink(ctx, request.GetUrl())
	if err != nil {
		logger.Error("GetShortLink", zap.Error(err))
		err = s.CheckError(err)
		return &pb.GetShortLinkResponse{Url: "{}"}, err
	}
	return &pb.GetShortLinkResponse{Url: url.ShortURL}, nil
}

func (s *Server) GetBaseLink(ctx context.Context, request *pb.GetBaseLinkRequest) (*pb.GetBaseLinkResponse, error) {
	logger.Info("GetBaseLink gRPC request", zap.String("url", request.GetUrl()))
	url, err := s.services.GetBaseLink(ctx, request.GetUrl())
	if err != nil {
		logger.Error("GetShortLink", zap.Error(err))
		err = s.CheckError(err)
		return &pb.GetBaseLinkResponse{Url: "{}"}, err
	}
	return &pb.GetBaseLinkResponse{Url: url.BaseURL}, nil
}

func (s *Server) CheckError(err error) error {
	if errors.Is(err, apperror.ErrNotFound) {
		return apperror.ErrGrpcNotFound
	}
	if errors.Is(err, apperror.ErrIsEmpty) {
		return apperror.ErrGrpcIsEmpty
	}
	return apperror.ErrGrpcServerError
}
