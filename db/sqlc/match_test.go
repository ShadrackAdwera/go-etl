package db

import (
	"context"
	"strconv"
	"testing"
	"time"

	"github.com/ShadrackAdwera/go-etl/utils"
	"github.com/stretchr/testify/require"
)

func CreateMatch(t *testing.T, user User) MatchDatum {
	file := CreateFile(t, user)

	newMatch := CreateMatchDataParams{
		HomeScored:  int32(utils.RandomInteger(0, 10)),
		AwayScored:  int32(utils.RandomInteger(0, 10)),
		HomeTeam:    utils.RandomString(6),
		AwayTeam:    utils.RandomString(5),
		MatchDate:   time.Now(),
		Referee:     utils.RandomString(7),
		Winner:      utils.RandomString(1),
		Season:      strconv.Itoa(time.Now().Year()),
		CreatedByID: user.ID,
		FileID:      file.ID,
	}

	match, err := testQuery.CreateMatchData(context.Background(), newMatch)

	require.NoError(t, err)
	require.NotEmpty(t, match)
	require.NotZero(t, match.ID)
	require.Equal(t, match.HomeScored, newMatch.HomeScored)
	require.Equal(t, match.AwayScored, newMatch.AwayScored)
	require.Equal(t, match.HomeTeam, newMatch.HomeTeam)
	require.Equal(t, match.AwayTeam, newMatch.AwayTeam)
	require.Equal(t, match.MatchDate, newMatch.MatchDate)
	require.Equal(t, match.Referee, newMatch.Referee)
	require.Equal(t, match.Winner, newMatch.Winner)
	require.Equal(t, match.Season, newMatch.Season)
	require.Equal(t, match.CreatedByID, newMatch.CreatedByID)
	require.Equal(t, match.FileID, newMatch.FileID)

	return match
}

func TestGetMatchData(t *testing.T) {
	n := 5

	user := CreateUser(t)

	for i := 0; i < n; i++ {
		CreateMatch(t, user)
	}

	matches, err := testQuery.GetMatchData(context.Background(), user.ID)

	require.NoError(t, err)
	require.NotEmpty(t, matches)
	require.Equal(t, len(matches), n)
}
