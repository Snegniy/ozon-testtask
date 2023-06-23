package service

import (
	"fmt"
	"github.com/Snegniy/ozon-testtask/internal/model"
)

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

type Repository interface {
	GetBaseURL(url string) (model.UrlShortStorage, error)
	GetShortURL(url string) (model.UrlBaseStorage, error)
	WriteNewLink(url, short string) (model.UrlBaseStorage, error)
}

func (s *Service) GetShortLink(url string) (model.UrlBaseStorage, error) {
	if url == "" {
		return model.UrlBaseStorage{}, fmt.Errorf("url cannot be empty")
	}
	res, err := s.repo.GetShortURL(url)
	if err != nil {
		short := s.GenerateLink()
		newLink, _ := s.repo.WriteNewLink(url, short)
		return newLink, nil
	}
	return res, err
}

func (s *Service) GetBaseLink(url string) (model.UrlShortStorage, error) {
	if url == "" {
		return model.UrlShortStorage{}, fmt.Errorf("url cannot be empty")
	}
	res, err := s.repo.GetBaseURL(url)
	return res, err
}
