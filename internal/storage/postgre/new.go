package postgre

import (
	"fmt"
	"loudy-back/configs/postgres"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Storage struct {
	db *pgxpool.Pool
}

const emptyValue = -1

func New() (*Storage, error) {
	pool, err := postgres.LoadPgxPool()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	return &Storage{db: pool}, nil
}
