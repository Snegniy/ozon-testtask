package postgre

import (
	"context"
	"fmt"
	"github.com/Snegniy/ozon-testtask/internal/config"
	"github.com/Snegniy/ozon-testtask/pkg/logger"
	"github.com/jackc/pgx/v4/pgxpool"
	"go.uber.org/zap"
)

const (
	newDB = `
		CREATE TABLE IF NOT EXISTS links 
		(
			shortlink  varchar(10)  PRIMARY KEY,
			baselink   text 	   NOT NULL UNIQUE
		);
	`
	table = "links"
)

type Repository struct {
	db *pgxpool.Pool
}

func NewRepository(config config.Config) (*Repository, error) {
	logger.Debug("Repo:Creating PostgresSQL repository")
	ctx := context.Background()
	var repo *Repository
	connString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", config.Postgres.Username, config.Postgres.Password, config.Postgres.Host, config.Postgres.Port, table)
	//connString := "postgres://admin:password@localhost:5432/links"
	pool, err := pgxpool.Connect(ctx, connString)
	if err != nil {
		logger.Error("not connected to base", zap.Error(err), zap.String("connection", connString))
		return nil, err
	}
	if err = pool.Ping(ctx); err != nil {
		logger.Error("not ping base", zap.Error(err))
		return nil, err
	}
	if err = repo.initDB(ctx); err != nil {
		logger.Error("fail init db", zap.Error(err))
		return nil, err
	}
	return repo, nil
}

func (r *Repository) initDB(ctx context.Context) error {
	logger.Debug("Repo:Init db")
	if _, err := r.db.Exec(ctx, newDB); err != nil {
		return err
	}
	return nil
}

func (r *Repository) GetBaseURL(url string) (string, error) {
	logger.Debug("Repo:Getting base URL from PostgresSQL storage", zap.String("url", url))
	var res string
	sql := fmt.Sprintf("SELECT shortlink FROM %s WHERE baselink = $1", table)
	if err := r.db.QueryRow(context.Background(), sql, url).Scan(&res); err != nil {
		logger.Warn("Couldn't find base URL", zap.String("shorturl", url))
		return "", err
	}
	return res, nil
}

func (r *Repository) GetShortURL(url string) (string, error) {
	logger.Debug("Repo:Getting short URL from PostgresSQL storage", zap.String("url", url))
	var res string
	sql := fmt.Sprintf("SELECT * FROM %s WHERE shortlink = $1", table)
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
