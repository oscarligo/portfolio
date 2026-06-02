package store

import (
	"backend/internal/repository"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Store struct {
	*repository.Queries
	db *pgxpool.Pool
}

func NewStore(pool *pgxpool.Pool) *Store {
	return &Store{
		Queries: repository.New(pool),
		db:      pool,
	}
}
