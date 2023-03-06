package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (srv *Server) getMatches(ctx *gin.Context) {
	// middleware
	matches, err := srv.store.GetMatchData(ctx, 1)

	if err != nil {
		ctx.JSON(http.StatusOK, errJSON(err))
		return
	}

	ctx.JSON(http.StatusOK, matches)
}
