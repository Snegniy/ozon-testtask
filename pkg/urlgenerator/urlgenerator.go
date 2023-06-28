package urlgenerator

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

func GenerateLink() string {
	logger.Debug("Generating link")
	res := make([]byte, 10)

	for i := range res {
		res[i] = chars[rand.Intn(len(chars))]
	}

	return string(res)
}
