package api

import (
	"encoding/csv"
	"fmt"
	"mime/multipart"
	"net/http"
	"strconv"
	"time"

	db "github.com/ShadrackAdwera/go-etl/db/sqlc"
	"github.com/ShadrackAdwera/go-etl/token"
	"github.com/ShadrackAdwera/go-etl/workers"
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

	// TODO: PUT THIS INSIDE A TRANSACTION
	createdFile, err := srv.store.CreateFile(ctx, db.CreateFileParams{
		FileUrl:     fmt.Sprintf("s3:://url:%s", matchData.MatchCsvFile.Filename),
		CreatedByID: user.ID,
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errJSON(err))
		return
	}

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
			FileID:      createdFile.ID,
		}

		matches = append(matches, matchDt)
	}

	err = srv.taskDistro.DistroSendFileDataToDb(ctx, &workers.DistroSendFileToDbPayload{
		Matches: matches,
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errJSON(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Data has been uploaded"})

}
