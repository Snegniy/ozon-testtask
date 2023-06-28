package service

import (
	"context"
	"fmt"
	"github.com/Snegniy/ozon-testtask/internal/model"
	"github.com/Snegniy/ozon-testtask/pkg/logger"
	"github.com/Snegniy/ozon-testtask/pkg/urlgenerator"
)

type Service struct {
	repo Repository
}

//go:generate mockgen -source=service.go -destination=mocks/mock.go

func NewService(repo Repository) *Service {
	logger.Debug("creating service")
	return &Service{repo: repo}
}

type Repository interface {
	GetBaseURL(ctx context.Context, url string) (string, error)
	GetShortURL(ctx context.Context, url string) (string, error)
	WriteNewLink(ctx context.Context, url, short string) (string, error)
}

func (s *Service) GetShortLink(ctx context.Context, url string) (model.UrlStorage, error) {
	logger.Debug("Service:GetShortLink")
	if url == "" {
		logger.Warn("empty url")
		return model.UrlStorage{}, fmt.Errorf("url cannot be empty")
	}
	res, err := s.repo.GetShortURL(ctx, url)
	if err != nil {
		var newLink string
		for {
			short := urlgenerator.GenerateLink()
			_, err := s.repo.GetBaseURL(ctx, short)
			if err != nil {
				newLink, _ = s.repo.WriteNewLink(ctx, url, short)
				break
			}
			logger.Warn("multiple links")
		}
		return model.UrlStorage{ShortURL: newLink}, nil
	}
	return model.UrlStorage{ShortURL: res}, err
}

func (s *Service) GetBaseLink(ctx context.Context, url string) (model.UrlStorage, error) {
	logger.Debug("Service:GetBaseLink")
	if url == "" {
		logger.Warn("empty url")
		return model.UrlStorage{}, fmt.Errorf("url cannot be empty")
	}
	res, err := s.repo.GetBaseURL(ctx, url)
	return model.UrlStorage{BaseURL: res}, err
}
