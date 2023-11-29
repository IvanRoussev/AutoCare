-- name: CreateMaintenance :one
INSERT INTO maintenances (
car_vin, maintenance_type, mileage
) VALUES (
$1, $2, $3
) RETURNING *;


-- name: GetMaintenanceByID :one
SELECT * FROM maintenances
WHERE maintenance_id = $1 LIMIT 1;


-- name: GetListMaintenancesByVIN :many
SELECT * FROM maintenances
WHERE car_vin = $1
LIMIT $2
OFFSET $3;


-- name: DeleteMaintenanceByVIN :exec
DELETE FROM maintenances WHERE car_vin = $1;

-- name: DeleteMaintenanceByID :exec
DELETE FROM maintenances WHERE maintenance_id = $1;