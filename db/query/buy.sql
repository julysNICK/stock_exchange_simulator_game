INSERT INTO buy (
  actionIdBuy,
  profileId,
  numberStocks,
  'limit'
) VALUES (
  $1,
  $2,
  $3,
  $4
) RETURNING *;

-- name: GetBuyById :one
SELECT * FROM buy WHERE id = $1;

-- name: GetBuyByProfileId :many
SELECT * FROM buy WHERE profileId = $1;

-- name: GetBuyByActionId :many
SELECT * FROM buy WHERE actionIdBuy = $1;

-- name: UpdateBuy :one
UPDATE buy SET
  actionIdBuy = COALESCE($1, actionIdBuy),
  profileId = COALESCE($2, profileId),
  numberStocks = COALESCE($3, numberStocks),
  'limit' = COALESCE($4, 'limit')
WHERE id = $5
RETURNING *;

-- name: DeleteBuy :one
DELETE FROM buy WHERE id = $1;
 
