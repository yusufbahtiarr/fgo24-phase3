package utils

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

func ConnectDB() (*pgxpool.Pool, error) {
	connStr := `postgres://postgres:222@localhost:5432/test3`

	pool, err := pgxpool.New(context.Background(), connStr)
	if err != nil {
		return nil, err
	}
	return pool, err

}
