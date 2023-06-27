package grpchandlers

import (
	"context"
	pb "github.com/Snegniy/ozon-testtask/pkg/api"
	"github.com/Snegniy/ozon-testtask/pkg/logger"
	"go.uber.org/zap"
)

func (s *Server) GetShortLink(ctx context.Context, request *pb.GetShortLinkRequest) (*pb.GetShortLinkResponse, error) {
	logger.Info("GetShortLink gRPC request", zap.String("url", request.GetUrl()))
	url, err := s.services.GetShortLink(ctx, request.GetUrl())
	if err != nil {
		logger.Error("GetShortLink", zap.Error(err))
		return &pb.GetShortLinkResponse{Url: "{}"}, err
	}
	return &pb.GetShortLinkResponse{Url: url.ShortURL}, nil
}

func (s *Server) GetBaseLink(ctx context.Context, request *pb.GetBaseLinkRequest) (*pb.GetBaseLinkResponse, error) {
	logger.Info("GetBaseLink gRPC request", zap.String("url", request.GetUrl()))
	url, err := s.services.GetBaseLink(ctx, request.GetUrl())
	if err != nil {
		logger.Error("GetShortLink", zap.Error(err))
		return &pb.GetBaseLinkResponse{Url: "{}"}, err
	}
	return &pb.GetBaseLinkResponse{Url: url.BaseURL}, nil
}
