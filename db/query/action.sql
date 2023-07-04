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

RETURNING *;

-- name: DeleteAction :exec
DELETE FROM actions WHERE id = $1;

-- name: ListActions :many
SELECT * FROM actions LIMIT $1 OFFSET $2;

-- name: CountActions :one
SELECT count(*) FROM actions;

  