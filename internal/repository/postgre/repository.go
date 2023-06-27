package postgre

import (
	"context"
	"fmt"
	"github.com/Snegniy/ozon-testtask/internal/config"
	"github.com/Snegniy/ozon-testtask/pkg/logger"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
	"time"
)

const (
	table = "links"
)

type Repository struct {
	db *pgxpool.Pool
}

func NewRepository(config config.Config) (*Repository, error) {
	logger.Debug("Repo:Creating PostgresSQL repository")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	pool, err := NewConnect(ctx, config)
	if err != nil {
		logger.Error("connect to DB error", zap.Error(err))
		return nil, err
	}
	return &Repository{db: pool}, nil

}

func NewConnect(ctx context.Context, config config.Config) (*pgxpool.Pool, error) {
	connString := fmt.Sprintf("postgres://%s:%s@linkshorter-db:%s/%s?sslmode=disable", config.Postgres.Username, config.Postgres.Password, config.Postgres.Port, table)
	pool, err := pgxpool.New(ctx, connString)
	if err != nil {
		logger.Error("not connected to base", zap.Error(err), zap.String("connection", connString))
		return nil, err
	}
	time.Sleep(time.Second * 2)
	if err = pool.Ping(ctx); err != nil {
		logger.Error("not ping base", zap.Error(err))
		return nil, err
	}
	return pool, nil
}

func (r *Repository) GetBaseURL(url string) (string, error) {
	logger.Debug("Repo:Getting base URL from PostgresSQL storage", zap.String("url", url))
	var res string
	sql := `SELECT baselink 
			FROM links 
			WHERE shortlink = $1`
	if err := r.db.QueryRow(context.Background(), sql, url).Scan(&res); err != nil {
		logger.Warn("Couldn't find base URL", zap.String("shorturl", url))
		return "", err
	}
	return res, nil
}

func (r *Repository) GetShortURL(url string) (string, error) {
	logger.Debug("Repo:Getting short URL from PostgresSQL storage", zap.String("url", url))
	var res string
	sql := `SELECT shortlink 
			FROM links 
			WHERE baselink = $1`
	if err := r.db.QueryRow(context.Background(), sql, url).Scan(&res); err != nil {
		logger.Warn("Couldn't find short URL", zap.String("baseurl", url))
		return "", err
	}
	return res, nil
}

func (r *Repository) WriteNewLink(url, short string) (string, error) {
	logger.Debug("Repo:Write new URL to PostgresSQL storage", zap.String("baseurl", url), zap.String("shorturl", short))
	if v, err := r.GetShortURL(url); err == nil {
		return v, nil
	}
	sql := fmt.Sprintf("INSERT INTO %s (shortlink,baselink) values ($1,$2)", table)
	if _, err := r.db.Exec(context.Background(), sql, short, url); err != nil {
		logger.Error("Couldn't insert new URL", zap.Error(err), zap.String("baseurl", url), zap.String("shorturl", short))
		return "", err
	}
	return short, nil
}
