package postgre

import (
	"context"
	"fmt"
	"github.com/Snegniy/ozon-testtask/internal/config"
	"github.com/jackc/pgx/v4/pgxpool"
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
	ctx := context.Background()
	var repo *Repository
	connString := fmt.Sprintf("postgres://%s:%s@%s/%s", config.Postgres.Username, config.Postgres.Password, config.Postgres.HostPort, table)
	pool, err := pgxpool.Connect(ctx, connString)
	if err != nil {
		return nil, err
	}
	if err = pool.Ping(ctx); err != nil {
		return nil, err
	}
	if err = repo.initDB(ctx); err != nil {
		return nil, err
	}
	return repo, nil
}

func (r *Repository) initDB(ctx context.Context) error {
	if _, err := r.db.Exec(ctx, newDB); err != nil {
		return err
	}
	return nil
}

func (r *Repository) GetBaseURL(url string) (string, error) {
	var res string
	sql := fmt.Sprintf("SELECT shortlink FROM %s WHERE baselink = $1", table)
	if err := r.db.QueryRow(context.Background(), sql, url).Scan(&res); err != nil {
		return "", err
	}
	return res, nil
}

func (r *Repository) GetShortURL(url string) (string, error) {
	var res string
	sql := fmt.Sprintf("SELECT * FROM %s WHERE shortlink = $1", table)
	if err := r.db.QueryRow(context.Background(), sql, url).Scan(&res); err != nil {
		return "", err
	}
	return res, nil
}

func (r *Repository) WriteNewLink(url, short string) (string, error) {
	if v, err := r.GetShortURL(url); err == nil {
		return v, nil
	}
	sql := fmt.Sprintf("INSERT INTO %s (shortlink,baselink) values ($1,$2)", table)
	if _, err := r.db.Exec(context.Background(), sql, short, url); err != nil {
		return "", err
	}
	return short, nil
}
