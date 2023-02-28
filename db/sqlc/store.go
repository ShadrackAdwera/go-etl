package db

import (
	"context"
	"database/sql"
)

// code for transactions go here
type TxStore interface {
	Querier
}

type Store struct {
	*Queries
	db *sql.DB
}

func NewStore(db *sql.DB) TxStore {
	return &Store{
		db:      db,
		Queries: New(db),
	}
}

func (store *Store) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, &sql.TxOptions{
		Isolation: sql.LevelReadCommitted,
	})

	if err != nil {
		return err
	}

	q := New(tx)
	err = fn(q)

	if err != nil {
		if txErr := tx.Rollback(); txErr != nil {
			return txErr
		}
		return err
	}
	return tx.Commit()
}
