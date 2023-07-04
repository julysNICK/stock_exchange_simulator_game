-- name: CreatePlayer :one
INSERT INTO players (
  username,
  hashed_password,
  full_name,
  cash,
  email
) values (
  $1,
  $2,
  $3,
  $4,
  $5
) 
RETURNING *;

-- name: GetPlayerByUsername :one
SELECT * FROM players WHERE username = $1;

-- name: GetPlayerByEmail :one
SELECT * FROM players WHERE email = $1;

-- name: GetPlayerById :one
SELECT * FROM players WHERE id_player = $1;

-- name: UpdatePlayer :one
UPDATE players SET
  username = COALESCE(sqlc.narg(username), username),
  hashed_password = COALESCE(sqlc.narg(hashed_password), hashed_password),
  full_name = COALESCE(sqlc.narg(full_name), full_name),
  cash = COALESCE(sqlc.narg(cash), cash),
  email = COALESCE(sqlc.narg(email), email),
  password_changed_at = COALESCE(sqlc.narg(password_changed_at), password_changed_at),
  created_at = COALESCE(sqlc.narg(created_at), created_at)
WHERE id_player = $1
RETURNING *;

-- name: DeletePlayer :exec
DELETE FROM players WHERE id_player = $1;

-- name: ListPlayers :many
SELECT * FROM players LIMIT $1 OFFSET $2;

-- name: CountPlayers :one
SELECT count(*) FROM players;

-- name: RankPlayers :many
SELECT * FROM players ORDER BY cash DESC LIMIT $1 OFFSET $2;
