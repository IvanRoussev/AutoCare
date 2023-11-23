package db

import (
	"context"
	"database/sql"
	"github.com/IvanRoussev/autocare/util"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func CreateRandomOwner(t *testing.T) Owner {
	arg := CreateOwnerParams{
		FirstName: util.RandomString(),
		LastName:  util.RandomString(),
		Country:   util.RandomString(),
	}

	owner, err := testQueries.CreateOwner(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, owner)

	require.Equal(t, arg.FirstName, owner.FirstName)
	require.Equal(t, arg.LastName, owner.LastName)
	require.Equal(t, arg.Country, owner.Country)

	require.NotZero(t, owner.ID)
	require.NotZero(t, owner.CreatedAt)
	return owner
}

func TestGetOwner(t *testing.T) {
	owner1 := CreateRandomOwner(t)
	owner2, err := testQueries.GetOwnerByID(context.Background(), owner1.ID)

	require.NoError(t, err)

	require.Equal(t, owner1.ID, owner2.ID)
	require.Equal(t, owner1.FirstName, owner2.FirstName)
	require.Equal(t, owner1.LastName, owner2.LastName)
	require.Equal(t, owner1.Country, owner2.Country)
	require.WithinDuration(t, owner1.CreatedAt, owner2.CreatedAt, time.Second)
}

func TestUpdate_OwnerFirstName_ByID(t *testing.T) {
	owner := CreateRandomOwner(t)

	params := UpdateOwnerFirstNameByIDParams{
		ID:        owner.ID,
		FirstName: "Jeremy",
	}

	ownerResponse, err := testQueries.UpdateOwnerFirstNameByID(context.Background(), params)

	require.NoError(t, err)
	require.NotEqual(t, owner.FirstName, ownerResponse.FirstName)
	require.Equal(t, ownerResponse.FirstName, params.FirstName)

}

func TestUpdate_OwnerLastName_ByID(t *testing.T) {
	owner := CreateRandomOwner(t)

	params := UpdateOwnerLastNameByIDParams{
		ID:       owner.ID,
		LastName: "Simon",
	}

	ownerResponse, err := testQueries.UpdateOwnerLastNameByID(context.Background(), params)

	require.NoError(t, err)
	require.NotEqual(t, owner.LastName, ownerResponse.LastName)
	require.Equal(t, ownerResponse.LastName, params.LastName)

}

func TestUpdate_OwnerCountry_ByID(t *testing.T) {
	owner := CreateRandomOwner(t)

	params := UpdateOwnerCountryByIDParams{
		ID:      owner.ID,
		Country: "Bulgaria",
	}

	ownerResponse, err := testQueries.UpdateOwnerCountryByID(context.Background(), params)

	require.NoError(t, err)
	require.NotEqual(t, owner.Country, ownerResponse.Country)
	require.Equal(t, ownerResponse.Country, params.Country)

}

func TestDeleteOwnerByID(t *testing.T) {
	owner := CreateRandomOwner(t)

	err := testQueries.DeleteOwnerByID(context.Background(), owner.ID)
	require.NoError(t, err)

	ownerResult, err := testQueries.GetOwnerByID(context.Background(), owner.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Zero(t, ownerResult)
}

func TestListOwners(t *testing.T) {
	for i := 0; i < 10; i++ {
		CreateRandomOwner(t)
	}

	arg := ListOwnersParams{
		Limit:  5,
		Offset: 5,
	}

	owners, err := testQueries.ListOwners(context.Background(), arg)
	require.NoError(t, err)
	require.Equal(t, len(owners), 5)

	for _, account := range owners {
		require.NotEmpty(t, account)
	}
}
