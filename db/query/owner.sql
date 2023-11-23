-- name: CreateOwner :one
INSERT INTO owner (
 first_name, last_name, country
) VALUES (
 $1, $2, $3
) RETURNING *;

-- name: GetOwnerByID :one
SELECT * FROM owner
WHERE id = $1 LIMIT 1;

-- name: ListOwners :many
SELECT * FROM owner
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateOwnerFirstNameByID :one
UPDATE owner
SET first_name = $2
WHERE id = $1
RETURNING  *;

-- name: UpdateOwnerLastNameByID :one
UPDATE owner
SET last_name = $2
WHERE id = $1
RETURNING *;

-- name: UpdateOwnerCountryByID :one
UPDATE owner
SET country = $2
WHERE id = $1
RETURNING *;

-- name: DeleteOwnerByID :exec
DELETE FROM owner WHERE id = $1;