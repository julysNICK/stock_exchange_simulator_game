-- name: CreateBuy :one
INSERT INTO buy (
  action_id_buy,
  profile_id,
  number_stocks,
  "limit_buy",
  "status"
) VALUES (
  $1,
  $2,
  $3,
  $4,
  $5
) RETURNING *;

-- name: GetBuyById :one
SELECT * FROM buy WHERE id = $1;

-- name: GetBuyByBuyIdAndProfileId :one
SELECT * FROM buy WHERE id = $1 AND profile_id = $2;

-- name: GetBuyByProfile_id :many
SELECT * FROM buy WHERE profile_id = $1;

-- name: GetBuyByActionId :many
SELECT * FROM buy WHERE action_id_buy = $1;

-- name: UpdateBuy :one
UPDATE buy SET
  action_id_buy = COALESCE(sqlc.narg(action_id_buy), action_id_buy),
  profile_id = COALESCE(sqlc.narg(profile_id), profile_id),
  number_stocks = COALESCE(sqlc.narg(number_stocks), number_stocks),
  "status" = COALESCE(sqlc.narg(status), "status"),
  limit_buy = COALESCE(sqlc.narg(limit_buy), limit_buy)
WHERE id = $1
RETURNING *;

-- name: DeleteBuy :one
DELETE FROM buy WHERE id = $1 RETURNING *;

-- name: ListBuy :many
SELECT * FROM buy LIMIT $1 OFFSET $2;

-- name: CountBuy :one
SELECT count(*) FROM buy;

-- name: ListBuyByProfile_id :many
SELECT * FROM buy WHERE profile_id = $1 LIMIT $2 OFFSET $3;
 
