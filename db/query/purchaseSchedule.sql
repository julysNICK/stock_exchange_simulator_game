-- name: CreatePurchaseSchedule :one
INSERT INTO  "purchaseSchedule" (
  "buyId",
  "stage",
  "created_order_buy"
) values (
  $1,
  $2,
  $3
)
RETURNING *;

-- name: GetAllPurchaseSchedule :many
SELECT * FROM "purchaseSchedule";

-- name: GetPurchaseScheduleById :one
SELECT * FROM "purchaseSchedule" WHERE id = $1;


-- name: UpdatePurchaseSchedule :one
UPDATE "purchaseSchedule" SET
  "buyId" = COALESCE($1, "buyId"),
  "stage" = COALESCE($2, "stage"),
  "created_order_buy" = COALESCE($3, "created_order_buy")
WHERE id = $4
RETURNING *;