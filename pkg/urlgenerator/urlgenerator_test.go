package urlgenerator

import (
	"github.com/Snegniy/ozon-testtask/pkg/logger"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGenerateLink(t *testing.T) {
	logger.Init("no")
	link := GenerateLink()

	require.NotNil(t, link)
	require.NotEmpty(t, link)
	require.Equal(t, 10, len(link))
}
