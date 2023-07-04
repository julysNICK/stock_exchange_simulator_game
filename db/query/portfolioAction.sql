-- name: createPortfolioAction :one
INSERT INTO portfolioActions (
  portfolioId,
  actionId,
  playerId,
  quantity,
  purchasePrice
) VALUES (
  $1,
  $2,
  $3,
  $4,
  $5
) RETURNING *;

-- name: getPortfolioActionById :one
SELECT * FROM portfolioActions WHERE id = $1;

-- name: getPortfolioActionByPortfolioId :many
SELECT * FROM portfolioActions WHERE portfolioId = $1;

-- name: getPortfolioActionByActionId :many
SELECT * FROM portfolioActions WHERE actionId = $1;

-- name: getPortfolioActionByPlayerId :many
SELECT * FROM portfolioActions WHERE playerId = $1;

-- name: updatePortfolioAction :one
UPDATE portfolioActions SET
  portfolioId = COALESCE($1, portfolioId),
  actionId = COALESCE($2, actionId),
  playerId = COALESCE($3, playerId),
  quantity = COALESCE($4, quantity),
  purchasePrice = COALESCE($5, purchasePrice)
WHERE id = $6
RETURNING *;

-- name: deletePortfolioAction :one
DELETE FROM portfolioActions WHERE id = $1;

-- name: listPortfolioActions :many
SELECT * FROM portfolioActions LIMIT $1 OFFSET $2;

-- name: countPortfolioActions :one
SELECT count(*) FROM portfolioActions;

-- name: listPortfolioActionsByPortfolioId :many
SELECT * FROM portfolioActions WHERE portfolioId = $1 LIMIT $2 OFFSET $3;
