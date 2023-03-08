package api

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	db "github.com/ShadrackAdwera/go-etl/db/sqlc"
	"github.com/ShadrackAdwera/go-etl/token"
	"github.com/ShadrackAdwera/go-etl/utils"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

func RandomUser(t *testing.T) db.User {
	username := utils.RandomString(10)

	hashedPw, err := utils.HashPassword(utils.RandomString(12))

	require.NoError(t, err)

	return db.User{
		ID:                utils.RandomInteger(1, 100),
		Username:          username,
		Email:             fmt.Sprintf("%s@mail.com", username),
		Password:          hashedPw,
		PasswordChangedAt: time.Now(),
		CreatedAt:         time.Now(),
	}
}

func setUpToken(t *testing.T,
	request *http.Request,
	tokenMaker token.TokenMaker,
	authorizationKey string,
	authTypeBearer string,
	username string,
	id int64, email string,
	duration time.Duration) {
	payload, tkn, err := tokenMaker.CreateToken(username, id, email, duration)

	require.NoError(t, err)
	require.NotEmpty(t, payload)

	tknHeader := fmt.Sprintf("%s %s", authTypeBearer, tkn)

	request.Header.Set(authorizationKey, tknHeader)
}

/*
1. Test OK
2. Test No Auth
3. Test Invalid Header
4. Test unsupported auth type
5. Test Expired token
*/

func TestMiddleware(t *testing.T) {

	testCases := []struct {
		name       string
		setUpAuth  func(t *testing.T, request *http.Request, tokenMaker token.TokenMaker)
		comparator func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name: "TestOK",
			setUpAuth: func(t *testing.T, request *http.Request, tokenMaker token.TokenMaker) {
				user := RandomUser(t)
				setUpToken(t, request, tokenMaker, authorization, authType, user.Username, user.ID, user.Email, time.Minute)
			},
			comparator: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
			},
		},
		{
			name: "TestNoAuth",
			setUpAuth: func(t *testing.T, request *http.Request, tokenMaker token.TokenMaker) {
			},
			comparator: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusUnauthorized, recorder.Code)
			},
		},
		{
			name: "TestInvalidHeader",
			setUpAuth: func(t *testing.T, request *http.Request, tokenMaker token.TokenMaker) {
				user := RandomUser(t)
				setUpToken(t, request, tokenMaker, "something", authType, user.Username, user.ID, user.Email, time.Minute)
			},
			comparator: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusUnauthorized, recorder.Code)
			},
		},
		{
			name: "TestUnsupportedAuthType",
			setUpAuth: func(t *testing.T, request *http.Request, tokenMaker token.TokenMaker) {
				user := RandomUser(t)
				setUpToken(t, request, tokenMaker, authorization, "unsupported", user.Username, user.ID, user.Email, time.Minute)
			},
			comparator: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusUnauthorized, recorder.Code)
			},
		},
		{
			name: "TestExpiredToken",
			setUpAuth: func(t *testing.T, request *http.Request, tokenMaker token.TokenMaker) {
				user := RandomUser(t)
				setUpToken(t, request, tokenMaker, authorization, authType, user.Username, user.ID, user.Email, -time.Minute)
			},
			comparator: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusUnauthorized, recorder.Code)
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			srv := newTestServer(t, nil)

			protectedRoute := "/protected-route"

			srv.router.GET(protectedRoute, authMiddleware(srv.tokenMaker), func(ctx *gin.Context) {
				ctx.JSON(http.StatusOK, gin.H{"message": "authenticated route handler"})
			})

			req, err := http.NewRequest(http.MethodGet, protectedRoute, nil)

			require.NoError(t, err)

			testCase.setUpAuth(t, req, srv.tokenMaker)

			recorder := httptest.NewRecorder()

			srv.router.ServeHTTP(recorder, req)
			testCase.comparator(t, recorder)
		})
	}
}
