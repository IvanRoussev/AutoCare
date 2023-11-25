-- name: CreateCar :one
INSERT INTO cars (vin, owner_id, make, model, year
) VALUES (
$1, $2, $3, $4, $5
) RETURNING *;

-- name: GetCarByVIN :one
SELECT * FROM cars
WHERE vin = $1 LIMIT 1;

-- name: ListCars :many
SELECT * FROM cars
ORDER BY vin
LIMIT $1
OFFSET $2;

-- name: ListCarsByOwnerID :many
SELECT * FROM cars
WHERE owner_id = $1
LIMIT $2
OFFSET $3;

-- name: UpdateCarOwnerIdByVIN :one
UPDATE cars
SET owner_id = $2
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