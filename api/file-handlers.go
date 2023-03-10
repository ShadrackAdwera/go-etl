package api

import (
	"encoding/csv"
	"mime/multipart"
	"net/http"
	"strconv"
	"time"

	db "github.com/ShadrackAdwera/go-etl/db/sqlc"
	"github.com/ShadrackAdwera/go-etl/token"
	"github.com/gin-gonic/gin"
)

func (srv *Server) getFiles(ctx *gin.Context) {
	user := ctx.MustGet(authPayload).(*token.TokenPayload)
	files, err := srv.store.GetFiles(ctx, user.ID)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errJSON(err))
		return
	}

	ctx.JSON(http.StatusOK, files)
}

type csvMatchData struct {
	MatchCsvFile *multipart.FileHeader `form:"file" binding:"required"`
}

/*
home_scored,
    away_scored,
    home_team,
    away_team,
    match_date,
    referee,
    winner,
    season,
    created_by_id,
    file_id
*/

func (srv *Server) uploadCsvFile(ctx *gin.Context) {
	var matchData csvMatchData

	if err := ctx.ShouldBind(&matchData); err != nil {
		ctx.JSON(http.StatusBadRequest, errJSON(err))
		return
	}

	file, err := matchData.MatchCsvFile.Open()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errJSON(err))
		return
	}

	defer file.Close()

	matchRecords, err := csv.NewReader(file).ReadAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errJSON(err))
		return
	}
	user := ctx.MustGet(authPayload).(*token.TokenPayload)
	var matches []db.CreateMatchDataParams

	for _, row := range matchRecords {
		hmSc, err := strconv.Atoi(row[3])
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errJSON(err))
			return
		}
		awSc, err := strconv.Atoi(row[4])
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errJSON(err))
			return
		}
		t, err := time.Parse("02/02/2006", row[0])
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errJSON(err))
			return
		}

		matchDt := db.CreateMatchDataParams{
			HomeTeam:    row[1],
			AwayTeam:    row[2],
			HomeScored:  int32(hmSc),
			AwayScored:  int32(awSc),
			Referee:     row[6],
			MatchDate:   t,
			Winner:      row[5],
			Season:      strconv.Itoa(t.Year()),
			CreatedByID: user.ID,
			FileID:      1,
		}

		matches = append(matches, matchDt)
	}

	ctx.JSON(http.StatusOK, matches)

}
