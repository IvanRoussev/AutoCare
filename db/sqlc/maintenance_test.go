package db

import (
	"context"
	"database/sql"
	"github.com/IvanRoussev/autocare/util"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func createRandomMaintenance(t *testing.T, car Car) Maintenance {
	arg := CreateMaintenanceParams{
		CarVin:          car.Vin,
		MaintenanceType: util.RandomString(6),
		Mileage:         util.RandomInt32(0, 100),
	}

	maintenance, err := testQueries.CreateMaintenance(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, maintenance)

	require.NotZero(t, maintenance.MaintenanceID)
	require.Equal(t, arg.CarVin, maintenance.CarVin)
	require.Equal(t, arg.MaintenanceType, maintenance.MaintenanceType)
	require.Equal(t, arg.Mileage, maintenance.Mileage)
	require.NotZero(t, maintenance.CreatedAt)

	return maintenance
}

//func TestGet_Maintenance_ByVIN(t *testing.T) {
//	owner := CreateRandomOwner(t)
//	car := createRandomCar(t, owner)
//	maintenance := createRandomMaintenance(t, car)
//
//	maintenanceResult, err := testQueries.GetMaintenanceByVIN(context.Background(), maintenance.CarVin)
//	require.NoError(t, err)
//	require.NotZero(t, maintenanceResult)
//
//	require.Equal(t, maintenance.MaintenanceID, maintenanceResult.MaintenanceID)
//	require.Equal(t, maintenance.MaintenanceType, maintenanceResult.MaintenanceType)
//	require.Equal(t, maintenance.CarVin, maintenanceResult.CarVin)
//	require.Equal(t, maintenance.Mileage, maintenanceResult.Mileage)
//	require.WithinDuration(t, maintenance.CreatedAt, maintenanceResult.CreatedAt, time.Second)
//}

func TestGet_Maintenance_ByID(t *testing.T) {
	user := createRandomUser(t)
	car := createRandomCar(t, user)
	maintenance := createRandomMaintenance(t, car)

	maintenanceResult, err := testQueries.GetMaintenanceByID(context.Background(), maintenance.MaintenanceID)
	require.NoError(t, err)
	require.NotZero(t, maintenanceResult)

	require.Equal(t, maintenance.MaintenanceID, maintenanceResult.MaintenanceID)
	require.Equal(t, maintenance.MaintenanceType, maintenanceResult.MaintenanceType)
	require.Equal(t, maintenance.CarVin, maintenanceResult.CarVin)
	require.Equal(t, maintenance.Mileage, maintenanceResult.Mileage)
	require.WithinDuration(t, maintenance.CreatedAt, maintenanceResult.CreatedAt, time.Second)
}

func TestUpdate_MaintenanceMileage_ByVIN(t *testing.T) {
	owner := createRandomUser(t)
	car := createRandomCar(t, owner)
	maintenance := createRandomMaintenance(t, car)

	arg := UpdateMaintenanceMileageByVINParams{
		CarVin:  car.Vin,
		Mileage: util.RandomInt32(0, 100000),
	}

	maintenanceResult, err := testQueries.UpdateMaintenanceMileageByVIN(context.Background(), arg)
	require.NoError(t, err)
	require.NotZero(t, maintenance)

	require.Equal(t, arg.Mileage, maintenanceResult.Mileage)
	require.Equal(t, maintenance.CarVin, maintenanceResult.CarVin)
	require.Equal(t, maintenance.MaintenanceType, maintenanceResult.MaintenanceType)
	require.Equal(t, maintenance.MaintenanceID, maintenanceResult.MaintenanceID)
	require.WithinDuration(t, maintenance.CreatedAt, maintenanceResult.CreatedAt, time.Second)
}

func TestUpdate_MaintenanceType_ByVIN(t *testing.T) {
	user := createRandomUser(t)
	car := createRandomCar(t, user)
	maintenance := createRandomMaintenance(t, car)

	arg := UpdateMaintenanceTypeByVINParams{
		CarVin:          car.Vin,
		MaintenanceType: util.RandomString(6),
	}

	maintenanceResult, err := testQueries.UpdateMaintenanceTypeByVIN(context.Background(), arg)
	require.NoError(t, err)
	require.NotZero(t, maintenance)

	require.Equal(t, maintenance.Mileage, maintenanceResult.Mileage)
	require.Equal(t, maintenance.CarVin, maintenanceResult.CarVin)
	require.Equal(t, arg.MaintenanceType, maintenanceResult.MaintenanceType)
	require.Equal(t, maintenance.MaintenanceID, maintenanceResult.MaintenanceID)
	require.WithinDuration(t, maintenance.CreatedAt, maintenanceResult.CreatedAt, time.Second)
}

func TestListMaintenances(t *testing.T) {

	for i := 0; i < 10; i++ {
		user := createRandomUser(t)
		car := createRandomCar(t, user)
		createRandomMaintenance(t, car)
	}

	arg := ListMaintenancesParams{
		Limit:  5,
		Offset: 5,
	}

	maintenances, err := testQueries.ListMaintenances(context.Background(), arg)
	require.NoError(t, err)
	require.Equal(t, len(maintenances), 5)

	for _, maintenance := range maintenances {
		require.NotEmpty(t, maintenance)
	}
}

//func TestDelete_Maintenance_ByVIN(t *testing.T) {
//	owner := CreateRandomOwner(t)
//	car := createRandomCar(t, owner)
//	maintenance := createRandomMaintenance(t, car)
//
//	err := testQueries.DeleteMaintenanceByVIN(context.Background(), maintenance.CarVin)
//	require.NoError(t, err)
//
//	maintenanceResult, err := testQueries.GetMaintenanceByVIN(context.Background(), maintenance.CarVin)
//	require.Error(t, err)
//	require.EqualError(t, err, sql.ErrNoRows.Error())
//	require.Zero(t, maintenanceResult)
//}

func TestDelete_Maintenance_ByID(t *testing.T) {
	user := createRandomUser(t)
	car := createRandomCar(t, user)
	maintenance := createRandomMaintenance(t, car)

	err := testQueries.DeleteMaintenanceByID(context.Background(), maintenance.MaintenanceID)
	require.NoError(t, err)

	maintenanceResult, err := testQueries.GetMaintenanceByID(context.Background(), maintenance.MaintenanceID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Zero(t, maintenanceResult)
}
