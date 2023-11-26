package db

import (
	"context"
	"database/sql"
	"github.com/IvanRoussev/autocare/util"
	"github.com/stretchr/testify/require"
	"testing"
)

func createRandomCar(t *testing.T, owner Owner) Car {
	arg := CreateCarParams{
		Vin:     util.RandomString(),
		OwnerID: owner.ID,
		Make:    util.RandomString(),
		Model:   util.RandomString(),
		Year:    util.RandomYear(),
	}

	car, err := testQueries.CreateCar(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, car)

	require.Equal(t, arg.Vin, car.Vin)
	require.Equal(t, arg.Make, car.Make)
	require.Equal(t, arg.Model, car.Model)
	require.Equal(t, arg.OwnerID, car.OwnerID)
	require.Equal(t, arg.Year, car.Year)
	return car
}

func TestGetCar(t *testing.T) {
	owner := CreateRandomOwner(t)
	car := createRandomCar(t, owner)

	carResult, err := testQueries.GetCarByVIN(context.Background(), car.Vin)
	require.NoError(t, err)
	require.NotEmpty(t, car)

	require.Equal(t, carResult.Vin, car.Vin)
	require.Equal(t, carResult.Make, car.Make)
	require.Equal(t, carResult.Model, car.Model)
	require.Equal(t, carResult.OwnerID, car.OwnerID)
	require.Equal(t, carResult.Year, car.Year)
}

func TestUpdateCarMakeByVIN(t *testing.T) {
	owner := CreateRandomOwner(t)
	car := createRandomCar(t, owner)

	arg := UpdateCarMakeByVINParams{
		Vin:  car.Vin,
		Make: util.RandomString(),
	}

	carResult, err := testQueries.UpdateCarMakeByVIN(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, car)

	require.Equal(t, carResult.Vin, car.Vin)
	require.Equal(t, carResult.Make, arg.Make)
	require.Equal(t, carResult.Model, car.Model)
	require.Equal(t, carResult.OwnerID, car.OwnerID)
	require.Equal(t, carResult.Year, car.Year)
}

func TestUpdate_CarModel_ByVIN(t *testing.T) {
	owner := CreateRandomOwner(t)
	car := createRandomCar(t, owner)

	arg := UpdateCarModelByVINParams{
		Vin:   car.Vin,
		Model: util.RandomString(),
	}

	carResult, err := testQueries.UpdateCarModelByVIN(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, car)

	require.Equal(t, carResult.Vin, car.Vin)
	require.Equal(t, carResult.Make, car.Make)
	require.Equal(t, carResult.Model, arg.Model)
	require.Equal(t, carResult.OwnerID, car.OwnerID)
	require.Equal(t, carResult.Year, car.Year)
}

func TestUpdate_CarOwner_ByVIN(t *testing.T) {
	owner := CreateRandomOwner(t)
	updatedOwner := CreateRandomOwner(t)
	car := createRandomCar(t, owner)

	arg := UpdateCarOwnerIdByVINParams{
		Vin:     car.Vin,
		OwnerID: updatedOwner.ID,
	}

	carResult, err := testQueries.UpdateCarOwnerIdByVIN(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, car)

	require.Equal(t, carResult.Vin, car.Vin)
	require.Equal(t, carResult.Make, car.Make)
	require.Equal(t, carResult.Model, car.Model)
	require.Equal(t, carResult.OwnerID, arg.OwnerID)
	require.Equal(t, carResult.Year, car.Year)
}

func TestUpdate_CarYear_ByVIN(t *testing.T) {
	owner := CreateRandomOwner(t)
	car := createRandomCar(t, owner)

	arg := UpdateCarYearByVINParams{
		Vin:  car.Vin,
		Year: util.RandomYear(),
	}

	carResult, err := testQueries.UpdateCarYearByVIN(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, car)

	require.Equal(t, carResult.Vin, car.Vin)
	require.Equal(t, carResult.Make, car.Make)
	require.Equal(t, carResult.Model, car.Model)
	require.Equal(t, carResult.OwnerID, car.OwnerID)
	require.Equal(t, carResult.Year, arg.Year)
}

func TestListCars(t *testing.T) {

	for i := 0; i < 10; i++ {
		owner := CreateRandomOwner(t)
		createRandomCar(t, owner)
	}

	arg := ListCarsParams{
		Limit:  5,
		Offset: 5,
	}

	cars, err := testQueries.ListCars(context.Background(), arg)
	require.NoError(t, err)
	require.Equal(t, len(cars), 5)

	for _, car := range cars {
		require.NotEmpty(t, car)
	}
}

func TestDeleteCarByID(t *testing.T) {
	owner := CreateRandomOwner(t)
	car := createRandomCar(t, owner)

	err := testQueries.DeleteCarByVIN(context.Background(), car.Vin)
	require.NoError(t, err)

	carResult, err := testQueries.GetCarByVIN(context.Background(), car.Vin)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Zero(t, carResult)
}
