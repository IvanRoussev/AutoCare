-- name: CreateUser :one
INSERT INTO users (
username, hash_password, full_name, email, country
) VALUES (
 $1, $2, $3, $4, $5
) RETURNING *;


-- name: GetUserByUsername :one
SELECT * FROM users
WHERE username = $1
LIMIT 1;

-- name: GetUserByID :one
SELECT * FROM users
WHERE id = $1
LIMIT 1;


-- name: DeleteUserByUsername :exec
DELETE FROM users WHERE username = $1;