package db

import (
	"context"
	"database/sql"
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

func TestGetMatchDataById(t *testing.T) {
	user := CreateUser(t)

	newMatch := CreateMatch(t, user)

	match, err := testQuery.GetMatchDataById(context.Background(), newMatch.ID)

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
}

func TestUpdateMatch(t *testing.T) {
	user := CreateUser(t)

	newMatch := CreateMatch(t, user)
	homeScore := int32(utils.RandomInteger(2, 10))

	updatedMatch, err := testQuery.UpdateMatchData(context.Background(), UpdateMatchDataParams{
		ID: newMatch.ID,
		HomeScored: sql.NullInt32{
			Int32: homeScore,
			Valid: true,
		},
	})

	require.NoError(t, err)
	require.NotEmpty(t, updatedMatch)
	require.Equal(t, updatedMatch.HomeScored, homeScore)
	require.Equal(t, updatedMatch.ID, newMatch.ID)
	require.Equal(t, updatedMatch.AwayScored, newMatch.AwayScored)
	require.Equal(t, updatedMatch.HomeTeam, newMatch.HomeTeam)
	require.Equal(t, updatedMatch.AwayTeam, newMatch.AwayTeam)
	require.Equal(t, updatedMatch.MatchDate, newMatch.MatchDate)
	require.Equal(t, updatedMatch.Referee, newMatch.Referee)
	require.Equal(t, updatedMatch.Winner, newMatch.Winner)
	require.Equal(t, updatedMatch.Season, newMatch.Season)
	require.Equal(t, updatedMatch.CreatedByID, newMatch.CreatedByID)
	require.Equal(t, updatedMatch.FileID, newMatch.FileID)
}
