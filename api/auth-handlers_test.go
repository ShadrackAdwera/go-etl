package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"

	mockdb "github.com/ShadrackAdwera/go-etl/db/mocks"
	db "github.com/ShadrackAdwera/go-etl/db/sqlc"
	"github.com/ShadrackAdwera/go-etl/utils"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

type eqCreateUserParamsMatcher struct {
	arg      db.CreateUserParams
	password string
}

func (e eqCreateUserParamsMatcher) Matches(x interface{}) bool {
	arg, ok := x.(db.CreateUserParams)
	if !ok {
		return false
	}

	err := utils.IsPassword(e.password, arg.Password)
	if err != nil {
		return false
	}

	e.arg.Password = arg.Password
	return reflect.DeepEqual(e.arg, arg)
}

func (e eqCreateUserParamsMatcher) String() string {
	return fmt.Sprintf("matches arg %v and password %v", e.arg, e.password)
}

func EqCreateUserParams(arg db.CreateUserParams, password string) gomock.Matcher {
	return eqCreateUserParamsMatcher{arg, password}
}

func CreateRandomUser(t *testing.T) (db.User, string) {
	username := utils.RandomString(8)
	password := utils.RandomString(12)

	hashedPw, err := utils.HashPassword(password)

	require.NoError(t, err)

	user := db.User{
		ID:                utils.RandomInteger(1, 20),
		Username:          username,
		Email:             fmt.Sprintf("%s@mail.com", username),
		Password:          hashedPw,
		PasswordChangedAt: time.Now(),
		CreatedAt:         time.Now(),
	}
	return user, password
}

func TestSignUp(t *testing.T) {
	user, password := CreateRandomUser(t)

	testCases := []struct {
		name       string
		body       db.CreateUserParams
		buildStubs func(t *testing.T, store *mockdb.MockTxStore)
		comparator func(t *testing.T, recoreder *httptest.ResponseRecorder)
	}{
		{
			name: "TestSuccessfulSignUp",
			body: db.CreateUserParams{
				Username: user.Username,
				Email:    user.Email,
				Password: password,
			},
			buildStubs: func(t *testing.T, store *mockdb.MockTxStore) {
				store.EXPECT().FindUserByEmail(gomock.Any(), gomock.Eq(user.Email)).Times(1).Return(db.User{}, sql.ErrNoRows)
				store.EXPECT().CreateUser(gomock.Any(), EqCreateUserParams(db.CreateUserParams{
					Username: user.Username,
					Email:    user.Email,
					Password: password,
				}, password)).Times(1).Return(user, nil)
			},
			comparator: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusCreated, recorder.Code)
				compareRequests(t, recorder.Body, user)
			},
		},
		{
			name: "TestInvalidInputs",
			body: db.CreateUserParams{
				Username: "-",
				Email:    user.Email,
				Password: password,
			},
			buildStubs: func(t *testing.T, store *mockdb.MockTxStore) {
				store.EXPECT().CreateUser(gomock.Any(), EqCreateUserParams(db.CreateUserParams{
					Username: "-",
					Email:    user.Email,
					Password: password,
				}, password)).Times(0)
			},
			comparator: func(t *testing.T, recoreder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recoreder.Code)
			},
		},
		{
			name: "TestInternalServerErrorFindUser",
			body: db.CreateUserParams{
				Username: user.Username,
				Email:    user.Email,
				Password: password,
			},
			buildStubs: func(t *testing.T, store *mockdb.MockTxStore) {
				store.EXPECT().FindUserByEmail(gomock.Any(), gomock.Eq(user.Email)).Times(1).Return(db.User{}, sql.ErrConnDone)
			},
			comparator: func(t *testing.T, recoreder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recoreder.Code)
			},
		},
		{
			name: "TestInternalServerError",
			body: db.CreateUserParams{
				Username: user.Username,
				Email:    user.Email,
				Password: password,
			},
			buildStubs: func(t *testing.T, store *mockdb.MockTxStore) {
				store.EXPECT().FindUserByEmail(gomock.Any(), gomock.Eq(user.Email)).Times(1).Return(db.User{}, sql.ErrNoRows)
				store.EXPECT().CreateUser(gomock.Any(), EqCreateUserParams(db.CreateUserParams{
					Username: user.Username,
					Email:    user.Email,
					Password: password,
				}, password)).Times(1).Return(db.User{}, sql.ErrConnDone)
			},
			comparator: func(t *testing.T, recoreder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recoreder.Code)
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			ctlr := gomock.NewController(t)

			store := mockdb.NewMockTxStore(ctlr)

			testCase.buildStubs(t, store)

			dt, err := json.Marshal(testCase.body)

			require.NoError(t, err)

			req, err := http.NewRequest(http.MethodPost, "/api/sign-up", bytes.NewReader(dt))

			require.NoError(t, err)

			srv := newTestServer(t, store)

			recorder := httptest.NewRecorder()

			srv.router.ServeHTTP(recorder, req)

			testCase.comparator(t, recorder)
		})
	}
}

func compareRequests(t *testing.T, responseBody *bytes.Buffer, user db.User) {
	var authResponse AuthResponse

	reader, err := io.ReadAll(responseBody)

	require.NoError(t, err)

	err = json.Unmarshal(reader, &authResponse)

	require.NoError(t, err)
	require.NotEmpty(t, authResponse)
	require.Equal(t, authResponse.User.ID, user.ID)
	require.Equal(t, authResponse.User.Username, user.Username)
	require.Equal(t, authResponse.User.Email, user.Email)
}
