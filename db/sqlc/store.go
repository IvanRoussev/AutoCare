package db

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/golang/mock/mockgen/model"
)

type Store interface {
	Querier
}

// SQLStore provides all features to execute SQL queries and transactions
type SQLStore struct {
	*Queries
	db *sql.DB
}

// NewStore creates a new Store
func NewStore(db *sql.DB) Store {
	return &SQLStore{
		db:      db,
		Queries: New(db),
	}
}

// ExecTx executes a function within a database transaction
func (store *SQLStore) execTx(ctx context.Context, fn func(queries *Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q := New(tx)
	err = fn(q)

	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, Rollback err: %v", err, rbErr)
		}
		return err
	}
	return tx.Commit()
}

//type MaintenanceTxParams struct {
//	CarVin          string `json:"car_vin"`
//	MaintenanceType string `json:"maintenance_type"`
//	Mileage         int32  `json:"mileage"`
//}
//
//type MaintenanceTxResult {
//
//}
//
//func (store *Store) MaintenanceTx(ctx context.Context, arg MaintenanceTxParams) MaintenanceTxResult {
//
//}
