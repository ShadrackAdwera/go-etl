package api

import (
	"net/http"

	"github.com/ShadrackAdwera/go-etl/token"
	"github.com/gin-gonic/gin"
)

func (srv *Server) getMatches(ctx *gin.Context) {
	user := ctx.MustGet(authPayload).(*token.TokenPayload)

	matches, err := srv.store.GetMatchData(ctx, user.ID)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errJSON(err))
		return
	}

	ctx.JSON(http.StatusOK, matches)
}
