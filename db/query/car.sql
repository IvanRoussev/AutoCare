-- name: CreateCar :one
INSERT INTO car (vin, owner_id, make, model, year
) VALUES (
$1, $2, $3, $4, $5
) RETURNING *;

-- name: GetCarByVIN :one
SELECT * FROM car
WHERE vin = $1 LIMIT 1;

-- name: ListCars :many
SELECT * FROM car
ORDER BY vin
LIMIT $1
OFFSET $2;

-- name: UpdateCarOwnerIdByVIN :one
UPDATE car
SET owner_id = $2
WHERE vin = $1
RETURNING  *;

-- name: UpdateCarMakeByVIN :one
UPDATE car
SET make = $2
WHERE vin = $1
RETURNING *;

-- name: UpdateCarModelByVIN :one
UPDATE car
SET model = $2
WHERE vin = $1
RETURNING *;

-- name: UpdateCarYearByVIN :one
UPDATE car
SET year = $2
WHERE vin = $1
RETURNING *;

-- name: DeleteCarByVIN :exec
DELETE FROM car WHERE vin = $1;