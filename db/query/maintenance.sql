-- name: CreateMaintenance :one
INSERT INTO maintenance (
car_vin, maintenance_type, mileage
) VALUES (
$1, $2, $3
) RETURNING *;

-- name: GetMaintenanceByVIN :one
SELECT * FROM maintenance
WHERE car_vin = $1 LIMIT 1;

-- name: GetMaintenanceByID :one
SELECT * FROM maintenance
WHERE maintenance_id = $1 LIMIT 1;

-- name: ListMaintenances :many
SELECT * FROM maintenance
ORDER BY car_vin
LIMIT $1
OFFSET $2;

-- name: UpdateMaintenanceTypeByVIN :one
UPDATE maintenance
SET maintenance_type = $2
WHERE car_vin = $1
RETURNING  *;

-- name: UpdateMaintenanceMileageByVIN :one
UPDATE maintenance
SET mileage = $2
WHERE car_vin = $1
RETURNING *;

-- name: DeleteMaintenanceByVIN :exec
DELETE FROM maintenance WHERE car_vin = $1;

-- name: DeleteMaintenanceByID :exec
DELETE FROM maintenance WHERE maintenance_id = $1;