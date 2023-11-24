-- name: CreateOwner :one
INSERT INTO owners (
 first_name, last_name, country
) VALUES (
 $1, $2, $3
) RETURNING *;

-- name: GetOwnerByID :one
SELECT * FROM owners
WHERE id = $1 LIMIT 1;

-- name: ListOwners :many
SELECT * FROM owners
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateOwnerFirstNameByID :one
UPDATE owners
SET first_name = $2
WHERE id = $1
RETURNING  *;

-- name: UpdateOwnerLastNameByID :one
UPDATE owners
SET last_name = $2
WHERE id = $1
RETURNING *;

-- name: UpdateOwnerCountryByID :one
UPDATE owners
SET country = $2
WHERE id = $1
RETURNING *;

-- name: DeleteOwnerByID :exec
DELETE FROM owners WHERE id = $1;