package service

import "math/rand"

const (
	charsLower = "abcdefghijklmnopqrstuvwxyz"
	charsUpper = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	digits     = "0123456789"
	special    = "_"
	chars      = charsLower + charsUpper + digits + special
)

func (s *Service) GenerateLink() string {
	res := make([]byte, 10)

	for {
		for i := range res {
			res[i] = chars[rand.Intn(len(chars))]
		}
		_, err := s.repo.GetBaseURL(string(res))

		if err != nil {
			break
		}
	}
	return string(res)
}
