-- name: CreatePortfolio :one
INSERT INTO portfolio (
  player_id
) VALUES (
  $1
) RETURNING *;

-- name: GetPortfolioById :one
SELECT * FROM portfolio WHERE id = $1;

-- name: GetPortfolioByPlayerId :one
SELECT * FROM portfolio WHERE player_id = $1;

-- name: UpdatePortfolio :one
UPDATE portfolio SET
  player_id = COALESCE($1, player_id)
WHERE id = $2
RETURNING *;

-- name: DeletePortfolio :one
DELETE FROM portfolio WHERE id = $1;

-- name: ListPortfolio :many
SELECT * FROM portfolio LIMIT $1 OFFSET $2;

-- name: CountPortfolio :one
SELECT count(*) FROM portfolio;

