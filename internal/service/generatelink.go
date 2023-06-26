package service

import (
	"github.com/Snegniy/ozon-testtask/pkg/logger"
	"math/rand"
)

const (
	charsLower = "abcdefghijklmnopqrstuvwxyz"
	charsUpper = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	digits     = "0123456789"
	special    = "_"
	chars      = charsLower + charsUpper + digits + special
)

func (s *Service) GenerateLink() string {
	logger.Debug("Generating link")
	res := make([]byte, 10)

	for {
		for i := range res {
			res[i] = chars[rand.Intn(len(chars))]
		}
		_, err := s.repo.GetBaseURL(string(res))

		if err != nil {
			break
		}
		logger.Warn("multiple links")
	}
	return string(res)
}
