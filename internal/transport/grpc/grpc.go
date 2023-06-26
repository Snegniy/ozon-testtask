package grpchandlers

import (
	"context"
	pb "github.com/Snegniy/ozon-testtask/pkg/api"
)

func (s *Server) GetShortLink(ctx context.Context, request *pb.GetShortLinkRequest) (*pb.GetShortLinkResponse, error) {
	url, err := s.services.GetShortLink(request.GetUrl())
	if err != nil {
		return &pb.GetShortLinkResponse{Url: "{}"}, err
	}

	return &pb.GetShortLinkResponse{Url: url.ShortURL}, nil
}

func (s *Server) GetBaseLink(ctx context.Context, request *pb.GetBaseLinkRequest) (*pb.GetBaseLinkResponse, error) {
	url, err := s.services.GetBaseLink(request.GetUrl())
	if err != nil {
		return &pb.GetBaseLinkResponse{Url: "{}"}, err
	}

	return &pb.GetBaseLinkResponse{Url: url.BaseURL}, nil
}
