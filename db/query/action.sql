-- name: CreateAction :one
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
RETURNING *;

-- name: GetActionByName :one
SELECT * FROM actions WHERE name = $1;

-- name: GetActionById :one
SELECT * FROM actions WHERE id = $1;

-- name: UpdateAction :one
UPDATE actions SET
  "name" = COALESCE(sqlc.narg(name), name),
  "id_actions" = COALESCE(sqlc.narg(id_actions), id_actions),
  "isin" = COALESCE(sqlc.narg(isin), isin),
  "wkn" = COALESCE(sqlc.narg(wkn), wkn),
  "current_value" = COALESCE(sqlc.narg(current_value), current_value),
  "bid" = COALESCE(sqlc.narg(bid), bid),
  "ask" = COALESCE(sqlc.narg(ask), ask),
  "spread" = COALESCE(sqlc.narg(spread), spread),
  "time_of_last_refresh" = COALESCE(sqlc.narg(time_of_last_refresh), time_of_last_refresh),
  "change_percentage" = COALESCE(sqlc.narg(change_percentage), change_percentage),
  "change_absolute" = COALESCE(sqlc.narg(change_absolute), change_absolute),
  peak24h = COALESCE(sqlc.narg(peak24h), peak24h),
  low24h = COALESCE(sqlc.narg(low24h), low24h),
  peak7d = COALESCE(sqlc.narg(peak7d), peak7d),
  low7d = COALESCE(sqlc.narg(low7d), low7d),
  peak30d = COALESCE(sqlc.narg(peak30d), peak30d),
  low30d = COALESCE(sqlc.narg(low30d), low30d),
  created_at = COALESCE(sqlc.narg(created_at), created_at)
WHERE id = $1

RETURNING *;

-- name: DeleteAction :exec
DELETE FROM actions WHERE id = $1;

-- name: ListActions :many
SELECT * FROM actions LIMIT $1 OFFSET $2;

-- name: CountActions :one
SELECT count(*) FROM actions;

  