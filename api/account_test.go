package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	mockdb "github.com/Coluding/udemy_backend/db/mock"
	db "github.com/Coluding/udemy_backend/db/sqlc"
	"github.com/Coluding/udemy_backend/util"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetAccountAPI(t *testing.T) {
	account := randomAccount()
	fmt.Println(account)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	store := mockdb.NewMockStore(ctrl)

	// we expect the GetAccount method to be called excatly one time, if not the test fails .Times(1)
	store.EXPECT().
		GetACCount(gomock.Any(), gomock.Eq(account.ID)).
		Times(1).
		Return(account, nil)

	server := NewServer(store)
	recorder := httptest.NewRecorder()

	url := fmt.Sprintf("/accounts/%d", account.ID)
	request, err := http.NewRequest(http.MethodGet, url, nil)
	require.NoError(t, err)

	server.router.ServeHTTP(recorder, request)
	require.Equal(t, http.StatusOK, recorder.Code)
	requireBodyMatchAccount(t, recorder.Body, account)
}

func randomAccount() db.Account {
	return db.Account{
		ID:       util.RandomInt(1, 1000),
		Owner:    util.RandomOwner(),
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}
}

func requireBodyMatchAccount(t *testing.T, body *bytes.Buffer, account db.Account) {
	data, err := ioutil.ReadAll(body)
	require.NoError(t, err)

	var gotAccount db.Account
	err = json.Unmarshal(data, &gotAccount)
	require.NoError(t, err)
	require.Equal(t, account, gotAccount)

}
