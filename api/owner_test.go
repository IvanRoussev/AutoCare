package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	mockdb "github.com/IvanRoussev/autocare/db/mock"
	db "github.com/IvanRoussev/autocare/db/sqlc"
	"github.com/IvanRoussev/autocare/util"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetOwnerAPI(t *testing.T) {
	owner := randomOwner()

	testCases := []struct {
		name          string
		ownerID       int64
		BuildStubs    func(store *mockdb.MockStore)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name:    "OK",
			ownerID: owner.ID,
			BuildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetOwnerByID(gomock.Any(), gomock.Eq(owner.ID)).
					Times(1).
					Return(owner, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchOwner(t, recorder.Body, owner)
			},
		},
	}
	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			store := mockdb.NewMockStore(ctrl)
			tc.BuildStubs(store)
			// Start test server and send request
			server := NewServer(store)
			recorder := httptest.NewRecorder()

			url := fmt.Sprintf("/owners/%d", owner.ID)
			request, err := http.NewRequest(http.MethodGet, url, nil)
			require.NoError(t, err)

			server.router.ServeHTTP(recorder, request)

			// Check response
			require.Equal(t, http.StatusOK, recorder.Code)
			requireBodyMatchOwner(t, recorder.Body, owner)
			tc.checkResponse(t, recorder)
		})
	}

}

func randomOwner() db.Owner {
	return db.Owner{
		ID:        util.RandomInt64(0, 1000),
		FirstName: util.RandomString(),
		LastName:  util.RandomString(),
		Country:   util.RandomString(),
	}
}

func requireBodyMatchOwner(t *testing.T, body *bytes.Buffer, owner db.Owner) {
	data, err := ioutil.ReadAll(body)
	require.NoError(t, err)

	var getOwner db.Owner
	err = json.Unmarshal(data, &getOwner)
	require.NoError(t, err)
	require.Equal(t, owner, getOwner)

}
