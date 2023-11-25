-- name: CreateMaintenance :one
INSERT INTO maintenances (
car_vin, maintenance_type, mileage
) VALUES (
$1, $2, $3
) RETURNING *;


-- name: GetMaintenanceByID :one
SELECT * FROM maintenances
WHERE maintenance_id = $1 LIMIT 1;

-- name: ListMaintenances :many
SELECT * FROM maintenances
ORDER BY car_vin
LIMIT $1
OFFSET $2;

-- name: GetListMaintenancesByVIN :many
SELECT * FROM maintenances
WHERE car_vin = $1
LIMIT $2
OFFSET $3;


-- name: UpdateMaintenanceTypeByVIN :one
UPDATE maintenances
SET maintenance_type = $2
WHERE car_vin = $1
RETURNING  *;

-- name: UpdateMaintenanceMileageByVIN :one
UPDATE maintenances
SET mileage = $2
WHERE car_vin = $1
RETURNING *;

-- name: DeleteMaintenanceByVIN :exec
DELETE FROM maintenances WHERE car_vin = $1;

-- name: DeleteMaintenanceByID :exec
DELETE FROM maintenances WHERE maintenance_id = $1;