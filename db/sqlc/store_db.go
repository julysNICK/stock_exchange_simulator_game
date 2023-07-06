package db

import (
	"context"
	"database/sql"
	"fmt"
)

type StoreDB interface {
	Querier
	BuyTx(ctx context.Context, arg BuyTxParams) (BuyTxResult, error)
	PlayerTx(ctx context.Context, arg PlayerTxParams) (PlayerTxResult, error)
}

type SQLStore struct {
	db *sql.DB
	*Queries
}

func NewStoreDB(db *sql.DB) StoreDB {
	return &SQLStore{
		db:      db,
		Queries: New(db), // this New() is for the queries we defined in sqlc
	}
}

func (store *SQLStore) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil) // nil means use default options

	if err != nil {
		return err
	}

	q := New(tx)

	err = fn(q)

	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx error: %v, rb error: %v", err, rbErr)
		}
		return err
	}

	return tx.Commit()
}
