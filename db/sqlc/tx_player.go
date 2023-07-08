package db

import (
	"context"
	"database/sql"
)

type PlayerTxParams struct {
	UserName       string `json:"username"`
	HashedPassword string `json:"hashed_password"`
	FullName       string `json:"full_name"`
	Cash           string `json:"cash"`
	Email          string `json:"email"`
}

type PlayerTxResult struct {
	Player    Player    `json:"player"`
	Portfolio Portfolio `json:"portfolio"`
}

func (store *SQLStore) PlayerTx(ctx context.Context, arg PlayerTxParams) (PlayerTxResult, error) {
	var result PlayerTxResult

	err := store.execTx(ctx, func(q *Queries) error {

		var err error

		result.Player, err = q.CreatePlayer(ctx, CreatePlayerParams{
			Username: sql.NullString{
				String: arg.UserName,
				Valid:  true,
			},
			HashedPassword: arg.HashedPassword,
			FullName:       arg.FullName,
			Cash:           arg.Cash,
			Email:          arg.Email,
		})

		if err != nil {
			return err
		}

		result.Portfolio, err = q.CreatePortfolio(ctx, result.Player.IDPlayer)

		if err != nil {
			return err
		}

		return nil
	})

	return result, err
}
