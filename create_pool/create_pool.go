package create_pool

import (
	"context"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

func CreatePool() *pgxpool.Pool {
	connStr := os.Getenv("CONN_DB")

	pool, err := pgxpool.New(context.Background(), connStr)
	if err != nil {
		panic(err)
	}

	return pool
}
