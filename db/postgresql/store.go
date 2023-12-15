package db

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

type Store struct {
	connPool *pgx.Conn
	*Queries
}

// NewStore creates a new store
func NewStore(connPool *pgx.Conn) *Store {
	return &Store{
		connPool: connPool,
		Queries:  New(connPool),
	}
}

func (store *Store) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.connPool.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return err
	}

	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(ctx); rbErr != nil {
			return fmt.Errorf("tx error")
		}
		return err
	}
	return tx.Commit(ctx)
}
