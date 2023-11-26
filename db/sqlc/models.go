// Code generated by sqlc. DO NOT EDIT.

package db

import (
	"time"
)

type Car struct {
	Vin     string `json:"vin"`
	OwnerID int64  `json:"owner_id"`
	Make    string `json:"make"`
	Model   string `json:"model"`
	Year    int64  `json:"year"`
}

type Maintenance struct {
	MaintenanceID   int32     `json:"maintenance_id"`
	CarVin          string    `json:"car_vin"`
	MaintenanceType string    `json:"maintenance_type"`
	Mileage         int32     `json:"mileage"`
	CreatedAt       time.Time `json:"created_at"`
}

type Owner struct {
	ID        int64     `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Country   string    `json:"country"`
	CreatedAt time.Time `json:"created_at"`
}
