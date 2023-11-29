-- name: CreateCar :one
INSERT INTO cars (vin, username, make, model, year
) VALUES (
$1, $2, $3, $4, $5
) RETURNING *;

-- name: GetCarByVIN :one
SELECT * FROM cars
WHERE vin = $1 LIMIT 1;


-- name: ListCarsByUsername :many
SELECT * FROM cars
WHERE username = $1
LIMIT $2
OFFSET $3;

-- name: UpdateCarUsernameByVIN :one
UPDATE cars
SET username = $2
WHERE vin = $1
RETURNING  *;

-- name: UpdateCarMakeByVIN :one
UPDATE cars
SET make = $2
WHERE vin = $1
RETURNING *;

-- name: UpdateCarModelByVIN :one
UPDATE cars
SET model = $2
WHERE vin = $1
RETURNING *;

-- name: UpdateCarYearByVIN :one
UPDATE cars
SET year = $2
WHERE vin = $1
RETURNING *;

-- name: DeleteCarByVIN :exec
DELETE FROM cars WHERE vin = $1;