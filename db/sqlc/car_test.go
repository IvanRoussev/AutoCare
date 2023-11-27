package db

import (
	"context"
	"database/sql"
	"github.com/IvanRoussev/autocare/util"
	"github.com/stretchr/testify/require"
	"testing"
)

func createRandomCar(t *testing.T, user User) Car {
	arg := CreateCarParams{
		Vin:    util.RandomString(6),
		UserID: user.ID,
		Make:   util.RandomString(6),
		Model:  util.RandomString(6),
		Year:   util.RandomYear(),
	}

	car, err := testQueries.CreateCar(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, car)

	require.Equal(t, arg.Vin, car.Vin)
	require.Equal(t, arg.Make, car.Make)
	require.Equal(t, arg.Model, car.Model)
	require.Equal(t, arg.UserID, car.UserID)
	require.Equal(t, arg.Year, car.Year)
	return car
}

func TestGetCar(t *testing.T) {
	user := createRandomUser(t)
	car := createRandomCar(t, user)

	carResult, err := testQueries.GetCarByVIN(context.Background(), car.Vin)
	require.NoError(t, err)
	require.NotEmpty(t, car)

	require.Equal(t, carResult.Vin, car.Vin)
	require.Equal(t, carResult.Make, car.Make)
	require.Equal(t, carResult.Model, car.Model)
	require.Equal(t, carResult.UserID, car.UserID)
	require.Equal(t, carResult.Year, car.Year)
}

func TestUpdateCarMakeByVIN(t *testing.T) {
	owner := createRandomUser(t)
	car := createRandomCar(t, owner)

	arg := UpdateCarMakeByVINParams{
		Vin:  car.Vin,
		Make: util.RandomString(6),
	}

	carResult, err := testQueries.UpdateCarMakeByVIN(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, car)

	require.Equal(t, carResult.Vin, car.Vin)
	require.Equal(t, carResult.Make, arg.Make)
	require.Equal(t, carResult.Model, car.Model)
	require.Equal(t, carResult.UserID, car.UserID)
	require.Equal(t, carResult.Year, car.Year)
}

func TestUpdate_CarModel_ByVIN(t *testing.T) {
	owner := createRandomUser(t)
	car := createRandomCar(t, owner)

	arg := UpdateCarModelByVINParams{
		Vin:   car.Vin,
		Model: util.RandomString(6),
	}

	carResult, err := testQueries.UpdateCarModelByVIN(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, car)

	require.Equal(t, carResult.Vin, car.Vin)
	require.Equal(t, carResult.Make, car.Make)
	require.Equal(t, carResult.Model, arg.Model)
	require.Equal(t, carResult.UserID, car.UserID)
	require.Equal(t, carResult.Year, car.Year)
}

func TestUpdate_CarOwner_ByVIN(t *testing.T) {
	owner := createRandomUser(t)
	updatedOwner := createRandomUser(t)
	car := createRandomCar(t, owner)

	arg := UpdateCarUserIdByVINParams{
		Vin:    car.Vin,
		UserID: updatedOwner.ID,
	}

	carResult, err := testQueries.UpdateCarUserIdByVIN(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, car)

	require.Equal(t, carResult.Vin, car.Vin)
	require.Equal(t, carResult.Make, car.Make)
	require.Equal(t, carResult.Model, car.Model)
	require.Equal(t, carResult.UserID, arg.UserID)
	require.Equal(t, carResult.Year, car.Year)
}

func TestUpdate_CarYear_ByVIN(t *testing.T) {
	user := createRandomUser(t)
	car := createRandomCar(t, user)

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
	require.Equal(t, carResult.UserID, car.UserID)
	require.Equal(t, carResult.Year, arg.Year)
}

func TestListCars(t *testing.T) {

	for i := 0; i < 10; i++ {
		user := createRandomUser(t)
		createRandomCar(t, user)
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
	owner := createRandomUser(t)
	car := createRandomCar(t, owner)

	err := testQueries.DeleteCarByVIN(context.Background(), car.Vin)
	require.NoError(t, err)

	carResult, err := testQueries.GetCarByVIN(context.Background(), car.Vin)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Zero(t, carResult)
}
