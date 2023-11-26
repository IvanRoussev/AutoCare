
-- name: GetUserByID :one
SELECT * FROM users
WHERE id = $1
LIMIT 1;

-- name: ListUsers :many
SELECT * FROM users
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateUserFullNameByID :one
UPDATE users
SET full_name = $2
WHERE id = $1
RETURNING  *;

-- name: UpdateUsernameByID :one
UPDATE users
SET username = $2
WHERE id = $1
RETURNING *;

-- name: UpdateUserCountryByID :one
UPDATE users
SET country = $2
WHERE id = $1
RETURNING *;

-- name: DeleteUserByID :exec
DELETE FROM users WHERE id = $1;