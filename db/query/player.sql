-- name: CreateProduct :one
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
  username = COALESCE($1, username),
  hashed_password = COALESCE($2, hashed_password),
  full_name = COALESCE($3, full_name),
  cash = COALESCE($4, cash),
  email = COALESCE($5, email),
  password_changed_at = COALESCE($6, password_changed_at),
  created_at = COALESCE($7, created_at)
WHERE id_player = $8
RETURNING *;

-- name: DeletePlayer :exec
DELETE FROM players WHERE id_player = $1;

-- name: ListPlayers :many
SELECT * FROM players LIMIT $1 OFFSET $2;

-- name: CountPlayers :one
SELECT count(*) FROM players;

-- name: RankPlayers :many
SELECT * FROM players ORDER BY cash DESC LIMIT $1 OFFSET $2;
