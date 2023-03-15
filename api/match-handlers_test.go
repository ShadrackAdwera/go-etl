package api

import (
	"database/sql"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"

	mockdb "github.com/ShadrackAdwera/go-etl/db/mocks"
	db "github.com/ShadrackAdwera/go-etl/db/sqlc"
	"github.com/ShadrackAdwera/go-etl/token"
	"github.com/ShadrackAdwera/go-etl/utils"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func CreateRandomMatch(id int64) []db.MatchDatum {
	n := 5
	matches := []db.MatchDatum{}
	for i := 0; i < n; i++ {
		match := db.MatchDatum{
			ID:          utils.RandomInteger(1, 50),
			HomeScored:  int32(utils.RandomInteger(1, 10)),
			AwayScored:  int32(utils.RandomInteger(1, 10)),
			HomeTeam:    utils.RandomString(10),
			AwayTeam:    utils.RandomString(12),
			MatchDate:   time.Now(),
			Referee:     utils.RandomString(15),
			Winner:      utils.RandomWinner(),
			Season:      strconv.Itoa(time.Now().Year()),
			CreatedAt:   time.Now(),
			CreatedByID: id,
			FileID:      utils.RandomInteger(1, 5),
		}
		matches = append(matches, match)
	}

	return matches
}

func TestGetMatchData(t *testing.T) {
	user := RandomUser(t)
	matches := CreateRandomMatch(user.ID)

	testCases := []struct {
		name       string
		setUpAuth  func(t *testing.T, req *http.Request, tokenMaker token.TokenMaker)
		buildStubs func(t *testing.T, store *mockdb.MockTxStore)
		results    func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name: "TestGetOK",
			setUpAuth: func(t *testing.T, req *http.Request, tokenMaker token.TokenMaker) {
				setUpToken(t, req, tokenMaker, authorization, authType, user.Username, user.ID, user.Email, time.Minute)
			},
			buildStubs: func(t *testing.T, store *mockdb.MockTxStore) {
				store.EXPECT().GetMatchData(gomock.Any(), gomock.Eq(user.ID)).Times(1).Return(matches, nil)
			},
			results: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
			},
		},
		{
			name: "TestInternalServerError",
			setUpAuth: func(t *testing.T, req *http.Request, tokenMaker token.TokenMaker) {
				setUpToken(t, req, tokenMaker, authorization, authType, user.Username, user.ID, user.Email, time.Minute)
			},
			buildStubs: func(t *testing.T, store *mockdb.MockTxStore) {
				store.EXPECT().GetMatchData(gomock.Any(), gomock.Eq(user.ID)).Times(1).Return([]db.MatchDatum{}, sql.ErrConnDone)
			},
			results: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			ctlr := gomock.NewController(t)

			store := mockdb.NewMockTxStore(ctlr)

			srv := newTestServer(t, store)

			req, err := http.NewRequest(http.MethodGet, "/api/matches", nil)

			require.NoError(t, err)
			testCase.setUpAuth(t, req, srv.tokenMaker)
			testCase.buildStubs(t, store)

			recorder := httptest.NewRecorder()

			srv.router.ServeHTTP(recorder, req)

			testCase.results(t, recorder)
		})

	}
}
