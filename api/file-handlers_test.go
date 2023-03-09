package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	mockdb "github.com/ShadrackAdwera/go-etl/db/mocks"
	db "github.com/ShadrackAdwera/go-etl/db/sqlc"
	"github.com/ShadrackAdwera/go-etl/token"
	"github.com/ShadrackAdwera/go-etl/utils"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

// func setUpAuthToken(t *testing.T, request *http.Request, tokenMaker token.TokenMaker, username string, id int64, email string, duration time.Duration) {
// 	payload, tkn, err := tokenMaker.CreateToken(username, id, email, duration)
// 	require.NoError(t, err)
// 	require.NotEmpty(t, payload)

// 	bearerToken := fmt.Sprintf("%s %s", authType, tkn)
// 	request.Header.Set(authorization, bearerToken)
// }

func TestGetFiles(t *testing.T) {
	file := GenerateRandomFile()

	testCases := []struct {
		name       string
		file       db.File
		setUpAuth  func(t *testing.T, request *http.Request, tokenMaker token.TokenMaker)
		buildStubs func(store *mockdb.MockTxStore)
		comparator func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name: "Status OK",
			file: file,
			setUpAuth: func(t *testing.T, request *http.Request, tokenMaker token.TokenMaker) {
				username := utils.RandomString(6)
				email := fmt.Sprintf("%s@mail.com", username)
				setUpToken(t, request, tokenMaker, authorization, authType, username, file.CreatedByID, email, time.Minute)
			},
			buildStubs: func(store *mockdb.MockTxStore) {
				store.EXPECT().GetFiles(gomock.Any(), gomock.Eq(file.CreatedByID)).Times(1).Return([]db.File{
					file,
				}, nil)
			},
			comparator: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				compareGetFilesRequest(t, recorder.Body, []db.File{
					file,
				})
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			ctlr := gomock.NewController(t)
			defer ctlr.Finish()
			store := mockdb.NewMockTxStore(ctlr)

			testCase.buildStubs(store)

			server := newTestServer(t, store)

			recorder := httptest.NewRecorder()

			request, err := http.NewRequest(http.MethodGet, "/api/files", nil)

			require.NoError(t, err)

			testCase.setUpAuth(t, request, server.tokenMaker)

			server.router.ServeHTTP(recorder, request)
			testCase.comparator(t, recorder)

		})
	}
}

func compareGetFilesRequest(t *testing.T, recorder *bytes.Buffer, files []db.File) {
	dt, err := io.ReadAll(recorder)
	require.NoError(t, err)

	var filesFound []db.File

	err = json.Unmarshal(dt, &filesFound)
	require.NoError(t, err)
	require.Equal(t, len(filesFound), len(files))
}

func GenerateRandomFile() db.File {
	return db.File{
		ID:          utils.RandomInteger(1, 300),
		FileUrl:     "/test-file-url-in-s3",
		CreatedAt:   time.Now(),
		CreatedByID: utils.RandomInteger(1, 10),
	}
}
