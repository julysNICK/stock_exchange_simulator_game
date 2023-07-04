-- name: CreateAction :one
INSERT INTO actions (
  name,
  idActions,
  ISIN,
  WKN,
  currentValue,
  BID,
  ASK,
  spread,
  timeOfLastRefresh,
  changePercentage,
  changeAbsolute,
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
  name = COALESCE($1, name),
  idActions = COALESCE($2, idActions),
  ISIN = COALESCE($3, ISIN),
  WKN = COALESCE($4, WKN),
  currentValue = COALESCE($5, currentValue),
  BID = COALESCE($6, BID),
  ASK = COALESCE($7, ASK),
  spread = COALESCE($8, spread),
  timeOfLastRefresh = COALESCE($9, timeOfLastRefresh),
  changePercentage = COALESCE($10, changePercentage),
  changeAbsolute = COALESCE($11, changeAbsolute),
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
  