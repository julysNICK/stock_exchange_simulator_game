// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2
// source: portfolio.sql

package db

import (
	"context"
)

const countPortfolio = `-- name: CountPortfolio :one
SELECT count(*) FROM portfolio
`

func (q *Queries) CountPortfolio(ctx context.Context) (int64, error) {
	row := q.db.QueryRowContext(ctx, countPortfolio)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const createPortfolio = `-- name: CreatePortfolio :one
INSERT INTO portfolio (
  player_id
) VALUES (
  $1
) RETURNING id, player_id, created_at
`

func (q *Queries) CreatePortfolio(ctx context.Context, playerID int64) (Portfolio, error) {
	row := q.db.QueryRowContext(ctx, createPortfolio, playerID)
	var i Portfolio
	err := row.Scan(&i.ID, &i.PlayerID, &i.CreatedAt)
	return i, err
}

const deletePortfolio = `-- name: DeletePortfolio :one
DELETE FROM portfolio WHERE id = $1 RETURNING id, player_id, created_at
`

func (q *Queries) DeletePortfolio(ctx context.Context, id int64) (Portfolio, error) {
	row := q.db.QueryRowContext(ctx, deletePortfolio, id)
	var i Portfolio
	err := row.Scan(&i.ID, &i.PlayerID, &i.CreatedAt)
	return i, err
}

const getPortfolioById = `-- name: GetPortfolioById :one
SELECT id, player_id, created_at FROM portfolio WHERE id = $1
`

func (q *Queries) GetPortfolioById(ctx context.Context, id int64) (Portfolio, error) {
	row := q.db.QueryRowContext(ctx, getPortfolioById, id)
	var i Portfolio
	err := row.Scan(&i.ID, &i.PlayerID, &i.CreatedAt)
	return i, err
}

const getPortfolioByPlayerId = `-- name: GetPortfolioByPlayerId :one
SELECT id, player_id, created_at FROM portfolio WHERE player_id = $1
`

func (q *Queries) GetPortfolioByPlayerId(ctx context.Context, playerID int64) (Portfolio, error) {
	row := q.db.QueryRowContext(ctx, getPortfolioByPlayerId, playerID)
	var i Portfolio
	err := row.Scan(&i.ID, &i.PlayerID, &i.CreatedAt)
	return i, err
}

const listPortfolio = `-- name: ListPortfolio :many
SELECT id, player_id, created_at FROM portfolio LIMIT $1 OFFSET $2
`

type ListPortfolioParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListPortfolio(ctx context.Context, arg ListPortfolioParams) ([]Portfolio, error) {
	rows, err := q.db.QueryContext(ctx, listPortfolio, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Portfolio{}
	for rows.Next() {
		var i Portfolio
		if err := rows.Scan(&i.ID, &i.PlayerID, &i.CreatedAt); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updatePortfolio = `-- name: UpdatePortfolio :one
UPDATE portfolio SET
  player_id = COALESCE($1, player_id)
WHERE id = $2
RETURNING id, player_id, created_at
`

type UpdatePortfolioParams struct {
	PlayerID int64 `json:"playerID"`
	ID       int64 `json:"id"`
}

func (q *Queries) UpdatePortfolio(ctx context.Context, arg UpdatePortfolioParams) (Portfolio, error) {
	row := q.db.QueryRowContext(ctx, updatePortfolio, arg.PlayerID, arg.ID)
	var i Portfolio
	err := row.Scan(&i.ID, &i.PlayerID, &i.CreatedAt)
	return i, err
}
