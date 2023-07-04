-- name: createPortfolioAction :one
INSERT INTO "portfolioActions" (
  portfolio_id,
  action_id,
  player_id,
  quantity,
  purchase_price
) VALUES (
  $1,
  $2,
  $3,
  $4,
  $5
) RETURNING *;

-- name: getPortfolioActionById :one
SELECT * FROM "portfolioActions" WHERE id = $1;

-- name: getPortfolioActionByPortfolio_id :many
SELECT * FROM "portfolioActions" WHERE portfolio_id = $1;

-- name: getPortfolioActionByAction_id :many
SELECT * FROM "portfolioActions" WHERE action_id = $1;

-- name: getPortfolioActionByPlayer_id :many
SELECT * FROM "portfolioActions" WHERE player_id = $1;

-- name: updatePortfolioAction :one
UPDATE "portfolioActions" SET
  portfolio_id = COALESCE($1, portfolio_id),
  action_id = COALESCE($2, action_id),
  player_id = COALESCE($3, player_id),
  quantity = COALESCE($4, quantity),
  purchase_price = COALESCE($5, purchase_price)
WHERE id = $6
RETURNING *;

-- name: deletePortfolioAction :one
DELETE FROM "portfolioActions" WHERE id = $1 RETURNING *;

-- name: listPortfolioActions :many
SELECT * FROM "portfolioActions" LIMIT $1 OFFSET $2;

-- name: countPortfolioActions :one
SELECT count(*) FROM "portfolioActions";

-- name: listPortfolioActionsByPortfolio_id :many
SELECT * FROM "portfolioActions" WHERE portfolio_id = $1 LIMIT $2 OFFSET $3;
