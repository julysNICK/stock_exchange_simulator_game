INSERT INTO buy (
  action_id_buy,
  profile_id,
  number_stocks,
  "limit"
) VALUES (
  $1,
  $2,
  $3,
  $4
) RETURNING *;

-- name: GetBuyById :one
SELECT * FROM buy WHERE id = $1;

-- name: GetBuyByProfile_id :many
SELECT * FROM buy WHERE profile_id = $1;

-- name: GetBuyByActionId :many
SELECT * FROM buy WHERE action_id_buy = $1;

-- name: UpdateBuy :one
UPDATE buy SET
  action_id_buy = COALESCE($1, action_id_buy),
  profile_id = COALESCE($2, profile_id),
  number_stocks = COALESCE($3, number_stocks),
  "limit" = COALESCE($4, "limit")
WHERE id = $5
RETURNING *;

-- name: DeleteBuy :one
DELETE FROM buy WHERE id = $1 RETURNING *;

-- name: ListBuy :many
SELECT * FROM buy LIMIT $1 OFFSET $2;

-- name: CountBuy :one
SELECT count(*) FROM buy;

-- name: ListBuyByProfile_id :many
SELECT * FROM buy WHERE profile_id = $1 LIMIT $2 OFFSET $3;
 
