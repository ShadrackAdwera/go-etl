package db

import "database/sql"

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
