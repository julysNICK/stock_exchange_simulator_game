// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2
// source: action.sql

package db

import (
	"context"
	"database/sql"
	"time"
)

const countActions = `-- name: CountActions :one
SELECT count(*) FROM actions
`

func (q *Queries) CountActions(ctx context.Context) (int64, error) {
	row := q.db.QueryRowContext(ctx, countActions)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const createAction = `-- name: CreateAction :one
INSERT INTO actions (
  "name",
  "id_actions",
  "isin",
  "wkn",
  "current_value",
  "bid",
  "ask",
  "spread",
  "time_of_last_refresh",
  "change_percentage",
  "change_absolute",
  peak24h,
  low24h,
  peak7d,
  low7d,
  peak30d,
  low30d
) values (
  $1,
  $2,
  $3,
  $4,
  $5,
  $6,
  $7,
  $8,
  $9,
  $10,
  $11,
  $12,
  $13,
  $14,
  $15,
  $16,
  $17
)
RETURNING id, name, id_actions, isin, wkn, current_value, bid, ask, spread, time_of_last_refresh, change_percentage, change_absolute, peak24h, low24h, peak7d, low7d, peak30d, low30d, created_at
`

type CreateActionParams struct {
	Name              string        `json:"name"`
	IDActions         sql.NullInt32 `json:"idActions"`
	Isin              string        `json:"isin"`
	Wkn               string        `json:"wkn"`
	CurrentValue      string        `json:"currentValue"`
	Bid               string        `json:"bid"`
	Ask               string        `json:"ask"`
	Spread            string        `json:"spread"`
	TimeOfLastRefresh time.Time     `json:"timeOfLastRefresh"`
	ChangePercentage  string        `json:"changePercentage"`
	ChangeAbsolute    string        `json:"changeAbsolute"`
	Peak24h           string        `json:"peak24h"`
	Low24h            string        `json:"low24h"`
	Peak7d            string        `json:"peak7d"`
	Low7d             string        `json:"low7d"`
	Peak30d           string        `json:"peak30d"`
	Low30d            string        `json:"low30d"`
}

func (q *Queries) CreateAction(ctx context.Context, arg CreateActionParams) (Action, error) {
	row := q.db.QueryRowContext(ctx, createAction,
		arg.Name,
		arg.IDActions,
		arg.Isin,
		arg.Wkn,
		arg.CurrentValue,
		arg.Bid,
		arg.Ask,
		arg.Spread,
		arg.TimeOfLastRefresh,
		arg.ChangePercentage,
		arg.ChangeAbsolute,
		arg.Peak24h,
		arg.Low24h,
		arg.Peak7d,
		arg.Low7d,
		arg.Peak30d,
		arg.Low30d,
	)
	var i Action
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.IDActions,
		&i.Isin,
		&i.Wkn,
		&i.CurrentValue,
		&i.Bid,
		&i.Ask,
		&i.Spread,
		&i.TimeOfLastRefresh,
		&i.ChangePercentage,
		&i.ChangeAbsolute,
		&i.Peak24h,
		&i.Low24h,
		&i.Peak7d,
		&i.Low7d,
		&i.Peak30d,
		&i.Low30d,
		&i.CreatedAt,
	)
	return i, err
}

const deleteAction = `-- name: DeleteAction :exec
DELETE FROM actions WHERE id = $1
`

func (q *Queries) DeleteAction(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteAction, id)
	return err
}

const getActionById = `-- name: GetActionById :one
SELECT id, name, id_actions, isin, wkn, current_value, bid, ask, spread, time_of_last_refresh, change_percentage, change_absolute, peak24h, low24h, peak7d, low7d, peak30d, low30d, created_at FROM actions WHERE id = $1
`

func (q *Queries) GetActionById(ctx context.Context, id int64) (Action, error) {
	row := q.db.QueryRowContext(ctx, getActionById, id)
	var i Action
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.IDActions,
		&i.Isin,
		&i.Wkn,
		&i.CurrentValue,
		&i.Bid,
		&i.Ask,
		&i.Spread,
		&i.TimeOfLastRefresh,
		&i.ChangePercentage,
		&i.ChangeAbsolute,
		&i.Peak24h,
		&i.Low24h,
		&i.Peak7d,
		&i.Low7d,
		&i.Peak30d,
		&i.Low30d,
		&i.CreatedAt,
	)
	return i, err
}

const getActionByName = `-- name: GetActionByName :one
SELECT id, name, id_actions, isin, wkn, current_value, bid, ask, spread, time_of_last_refresh, change_percentage, change_absolute, peak24h, low24h, peak7d, low7d, peak30d, low30d, created_at FROM actions WHERE name = $1
`

func (q *Queries) GetActionByName(ctx context.Context, name string) (Action, error) {
	row := q.db.QueryRowContext(ctx, getActionByName, name)
	var i Action
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.IDActions,
		&i.Isin,
		&i.Wkn,
		&i.CurrentValue,
		&i.Bid,
		&i.Ask,
		&i.Spread,
		&i.TimeOfLastRefresh,
		&i.ChangePercentage,
		&i.ChangeAbsolute,
		&i.Peak24h,
		&i.Low24h,
		&i.Peak7d,
		&i.Low7d,
		&i.Peak30d,
		&i.Low30d,
		&i.CreatedAt,
	)
	return i, err
}

const listActions = `-- name: ListActions :many
SELECT id, name, id_actions, isin, wkn, current_value, bid, ask, spread, time_of_last_refresh, change_percentage, change_absolute, peak24h, low24h, peak7d, low7d, peak30d, low30d, created_at FROM actions LIMIT $1 OFFSET $2
`

type ListActionsParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListActions(ctx context.Context, arg ListActionsParams) ([]Action, error) {
	rows, err := q.db.QueryContext(ctx, listActions, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Action{}
	for rows.Next() {
		var i Action
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.IDActions,
			&i.Isin,
			&i.Wkn,
			&i.CurrentValue,
			&i.Bid,
			&i.Ask,
			&i.Spread,
			&i.TimeOfLastRefresh,
			&i.ChangePercentage,
			&i.ChangeAbsolute,
			&i.Peak24h,
			&i.Low24h,
			&i.Peak7d,
			&i.Low7d,
			&i.Peak30d,
			&i.Low30d,
			&i.CreatedAt,
		); err != nil {
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

const updateAction = `-- name: UpdateAction :one
UPDATE actions SET
  "name" = COALESCE($1, "name"),
  "id_actions" = COALESCE($2, "id_actions"),
  "isin" = COALESCE($3, "isin"),
  "wkn" = COALESCE($4, "wkn"),
  "current_value" = COALESCE($5, "current_value"),
  "bid" = COALESCE($6, "bid"),
  "ask" = COALESCE($7, "ask"),
  "spread" = COALESCE($8, "spread"),
  "time_of_last_refresh" = COALESCE($9, "time_of_last_refresh"),
  "change_percentage" = COALESCE($10, "change_percentage"),
  "change_absolute" = COALESCE($11, "change_absolute"),
  peak24h = COALESCE($12, peak24h),
  low24h = COALESCE($13, low24h),
  peak7d = COALESCE($14, peak7d),
  low7d = COALESCE($15, low7d),
  peak30d = COALESCE($16, peak30d),
  low30d = COALESCE($17, low30d),
  created_at = COALESCE($18, created_at)
WHERE id = $19

RETURNING id, name, id_actions, isin, wkn, current_value, bid, ask, spread, time_of_last_refresh, change_percentage, change_absolute, peak24h, low24h, peak7d, low7d, peak30d, low30d, created_at
`

type UpdateActionParams struct {
	Name              string        `json:"name"`
	IDActions         sql.NullInt32 `json:"idActions"`
	Isin              string        `json:"isin"`
	Wkn               string        `json:"wkn"`
	CurrentValue      string        `json:"currentValue"`
	Bid               string        `json:"bid"`
	Ask               string        `json:"ask"`
	Spread            string        `json:"spread"`
	TimeOfLastRefresh time.Time     `json:"timeOfLastRefresh"`
	ChangePercentage  string        `json:"changePercentage"`
	ChangeAbsolute    string        `json:"changeAbsolute"`
	Peak24h           string        `json:"peak24h"`
	Low24h            string        `json:"low24h"`
	Peak7d            string        `json:"peak7d"`
	Low7d             string        `json:"low7d"`
	Peak30d           string        `json:"peak30d"`
	Low30d            string        `json:"low30d"`
	CreatedAt         time.Time     `json:"createdAt"`
	ID                int64         `json:"id"`
}

func (q *Queries) UpdateAction(ctx context.Context, arg UpdateActionParams) (Action, error) {
	row := q.db.QueryRowContext(ctx, updateAction,
		arg.Name,
		arg.IDActions,
		arg.Isin,
		arg.Wkn,
		arg.CurrentValue,
		arg.Bid,
		arg.Ask,
		arg.Spread,
		arg.TimeOfLastRefresh,
		arg.ChangePercentage,
		arg.ChangeAbsolute,
		arg.Peak24h,
		arg.Low24h,
		arg.Peak7d,
		arg.Low7d,
		arg.Peak30d,
		arg.Low30d,
		arg.CreatedAt,
		arg.ID,
	)
	var i Action
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.IDActions,
		&i.Isin,
		&i.Wkn,
		&i.CurrentValue,
		&i.Bid,
		&i.Ask,
		&i.Spread,
		&i.TimeOfLastRefresh,
		&i.ChangePercentage,
		&i.ChangeAbsolute,
		&i.Peak24h,
		&i.Low24h,
		&i.Peak7d,
		&i.Low7d,
		&i.Peak30d,
		&i.Low30d,
		&i.CreatedAt,
	)
	return i, err
}
