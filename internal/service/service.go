package service

import (
	"fmt"
	"github.com/Snegniy/ozon-testtask/internal/model"
	"github.com/Snegniy/ozon-testtask/pkg/logger"
)

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	logger.Debug("creating service")
	return &Service{repo: repo}
}

type Repository interface {
	GetBaseURL(url string) (string, error)
	GetShortURL(url string) (string, error)
	WriteNewLink(url, short string) (string, error)
}

func (s *Service) GetShortLink(url string) (model.UrlStorage, error) {
	logger.Debug("Service:GetShortLink")
	if url == "" {
		logger.Warn("empty url")
		return model.UrlStorage{}, fmt.Errorf("url cannot be empty")
	}
	res, err := s.repo.GetShortURL(url)
	if err != nil {
		short := s.GenerateLink()
		newLink, _ := s.repo.WriteNewLink(url, short)
		return model.UrlStorage{ShortURL: newLink}, nil
	}
	return model.UrlStorage{ShortURL: res}, err
}

func (s *Service) GetBaseLink(url string) (model.UrlStorage, error) {
	logger.Debug("Service:GetBaseLink")
	if url == "" {
		logger.Warn("empty url")
		return model.UrlStorage{}, fmt.Errorf("url cannot be empty")
	}
	res, err := s.repo.GetBaseURL(url)
	return model.UrlStorage{BaseURL: res}, err
}
