package postgres

import (
	"context"
	"github.com/Snegniy/ozon-testtask/internal/config"
	"github.com/Snegniy/ozon-testtask/pkg/logger"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
	"time"
)

func NewConnect(ctx context.Context, config config.Config) (*pgxpool.Pool, error) {
	pool, err := pgxpool.New(ctx, config.Postgres.ConnString)
	if err != nil {
		logger.Error("not connected to base", zap.Error(err), zap.String("connection", config.Postgres.ConnString))
		return nil, err
	}
	time.Sleep(time.Second * 2)
	if err = pool.Ping(ctx); err != nil {
		logger.Error("not ping base", zap.Error(err))
		return nil, err
	}
	return pool, nil
}
