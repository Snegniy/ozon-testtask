package memdb

import (
	"context"
	"github.com/Snegniy/ozon-testtask/pkg/logger"
	"github.com/stretchr/testify/require"
	"testing"
)

var baseUrl = "ozon.ru"
var shortenUrl = "AAAAAAAAAA"
var ctx = context.Background()

func TestRepository_WriteNewLink(t *testing.T) {
	logger.Init("no")
	db := NewRepository()

	res, err := db.WriteNewLink(ctx, baseUrl, shortenUrl)

	require.Equal(t, shortenUrl, res)
	require.NoError(t, err)
}

func TestRepository_GetBaseURL(t *testing.T) {
	logger.Init("no")
	db := NewRepository()

	_, _ = db.WriteNewLink(ctx, baseUrl, shortenUrl)
	res, err := db.GetBaseURL(ctx, shortenUrl)

	require.Equal(t, baseUrl, res)
	require.NoError(t, err)
}

func TestRepository_GetBaseURL2(t *testing.T) {
	logger.Init("no")
	db := NewRepository()

	_, err := db.GetBaseURL(ctx, shortenUrl)

	require.Error(t, err)
}

func TestRepository_GetShortURL(t *testing.T) {
	logger.Init("no")
	db := NewRepository()

	_, _ = db.WriteNewLink(ctx, baseUrl, shortenUrl)
	res, err := db.GetShortURL(ctx, baseUrl)

	require.Equal(t, shortenUrl, res)
	require.NoError(t, err)
}
