package postgre

import (
	"context"
	"github.com/Snegniy/ozon-testtask/pkg/logger"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

type Repository struct {
	db *pgxpool.Pool
}

func NewRepository(con *pgxpool.Pool) (*Repository, error) {
	logger.Debug("Repo:Creating PostgresSQL repository")
	return &Repository{db: con}, nil
}

func (r *Repository) GetBaseURL(ctx context.Context, url string) (string, error) {
	logger.Debug("Repo:Getting base URL from PostgresSQL storage", zap.String("url", url))
	var res string
	sql := `SELECT baselink 
			FROM  links
			WHERE shortlink = $1`
	if err := r.db.QueryRow(ctx, sql, url).Scan(&res); err != nil {
		logger.Warn("Couldn't find base URL", zap.String("shorturl", url))
		return "", err
	}
	return res, nil
}

func (r *Repository) GetShortURL(ctx context.Context, url string) (string, error) {
	logger.Debug("Repo:Getting short URL from PostgresSQL storage", zap.String("url", url))
	var res string
	sql := `SELECT shortlink 
			FROM links 
			WHERE baselink = $1`
	if err := r.db.QueryRow(ctx, sql, url).Scan(&res); err != nil {
		logger.Warn("Couldn't find short URL", zap.String("baseurl", url))
		return "", err
	}
	return res, nil
}

func (r *Repository) WriteNewLink(ctx context.Context, url, short string) (string, error) {
	logger.Debug("Repo:Write new URL to PostgresSQL storage", zap.String("baseurl", url), zap.String("shorturl", short))
	if v, err := r.GetShortURL(ctx, url); err == nil {
		return v, nil
	}
	sql := "INSERT INTO links (shortlink,baselink) values ($1,$2)"
	if _, err := r.db.Exec(context.Background(), sql, short, url); err != nil {
		logger.Error("Couldn't insert new URL", zap.Error(err), zap.String("baseurl", url), zap.String("shorturl", short))
		return "", err
	}
	return short, nil
}
