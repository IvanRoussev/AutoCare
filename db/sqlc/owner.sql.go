// Code generated by sqlc. DO NOT EDIT.
// source: owner.sql

package db

import (
	"context"
)

const createOwner = `-- name: CreateOwner :one
INSERT INTO owner (
 first_name, last_name, country
) VALUES (
 $1, $2, $3
) RETURNING id, first_name, last_name, country, created_at
`

type CreateOwnerParams struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Country   string `json:"country"`
}

func (q *Queries) CreateOwner(ctx context.Context, arg CreateOwnerParams) (Owner, error) {
	row := q.db.QueryRowContext(ctx, createOwner, arg.FirstName, arg.LastName, arg.Country)
	var i Owner
	err := row.Scan(
		&i.ID,
		&i.FirstName,
		&i.LastName,
		&i.Country,
		&i.CreatedAt,
	)
	return i, err
}

const deleteOwnerByID = `-- name: DeleteOwnerByID :exec
DELETE FROM owner WHERE id = $1
`

func (q *Queries) DeleteOwnerByID(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteOwnerByID, id)
	return err
}

const getOwnerByID = `-- name: GetOwnerByID :one
SELECT id, first_name, last_name, country, created_at FROM owner
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetOwnerByID(ctx context.Context, id int64) (Owner, error) {
	row := q.db.QueryRowContext(ctx, getOwnerByID, id)
	var i Owner
	err := row.Scan(
		&i.ID,
		&i.FirstName,
		&i.LastName,
		&i.Country,
		&i.CreatedAt,
	)
	return i, err
}

const listOwners = `-- name: ListOwners :many
SELECT id, first_name, last_name, country, created_at FROM owner
ORDER BY id
LIMIT $1
OFFSET $2
`

type ListOwnersParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListOwners(ctx context.Context, arg ListOwnersParams) ([]Owner, error) {
	rows, err := q.db.QueryContext(ctx, listOwners, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Owner
	for rows.Next() {
		var i Owner
		if err := rows.Scan(
			&i.ID,
			&i.FirstName,
			&i.LastName,
			&i.Country,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateOwnerCountryByID = `-- name: UpdateOwnerCountryByID :one
UPDATE owner
SET country = $2
WHERE id = $1
RETURNING id, first_name, last_name, country, created_at
`

type UpdateOwnerCountryByIDParams struct {
	ID      int64  `json:"id"`
	Country string `json:"country"`
}

func (q *Queries) UpdateOwnerCountryByID(ctx context.Context, arg UpdateOwnerCountryByIDParams) (Owner, error) {
	row := q.db.QueryRowContext(ctx, updateOwnerCountryByID, arg.ID, arg.Country)
	var i Owner
	err := row.Scan(
		&i.ID,
		&i.FirstName,
		&i.LastName,
		&i.Country,
		&i.CreatedAt,
	)
	return i, err
}

const updateOwnerFirstNameByID = `-- name: UpdateOwnerFirstNameByID :one
UPDATE owner
SET first_name = $2
WHERE id = $1
RETURNING  id, first_name, last_name, country, created_at
`

type UpdateOwnerFirstNameByIDParams struct {
	ID        int64  `json:"id"`
	FirstName string `json:"first_name"`
}

func (q *Queries) UpdateOwnerFirstNameByID(ctx context.Context, arg UpdateOwnerFirstNameByIDParams) (Owner, error) {
	row := q.db.QueryRowContext(ctx, updateOwnerFirstNameByID, arg.ID, arg.FirstName)
	var i Owner
	err := row.Scan(
		&i.ID,
		&i.FirstName,
		&i.LastName,
		&i.Country,
		&i.CreatedAt,
	)
	return i, err
}

const updateOwnerLastNameByID = `-- name: UpdateOwnerLastNameByID :one
UPDATE owner
SET last_name = $2
WHERE id = $1
RETURNING id, first_name, last_name, country, created_at
`

type UpdateOwnerLastNameByIDParams struct {
	ID       int64  `json:"id"`
	LastName string `json:"last_name"`
}

func (q *Queries) UpdateOwnerLastNameByID(ctx context.Context, arg UpdateOwnerLastNameByIDParams) (Owner, error) {
	row := q.db.QueryRowContext(ctx, updateOwnerLastNameByID, arg.ID, arg.LastName)
	var i Owner
	err := row.Scan(
		&i.ID,
		&i.FirstName,
		&i.LastName,
		&i.Country,
		&i.CreatedAt,
	)
	return i, err
}
