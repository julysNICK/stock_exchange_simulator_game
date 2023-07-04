-- name: CreatePortfolio :one
INSERT INTO portfolio (
  playerId
) VALUES (
  $1
) RETURNING *;

-- name: GetPortfolioById :one
SELECT * FROM portfolio WHERE id = $1;

-- name: GetPortfolioByPlayerId :one
SELECT * FROM portfolio WHERE playerId = $1;

-- name: UpdatePortfolio :one
UPDATE portfolio SET
  playerId = COALESCE($1, playerId)
WHERE id = $2
RETURNING *;

-- name: DeletePortfolio :one
DELETE FROM portfolio WHERE id = $1;