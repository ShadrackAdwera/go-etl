package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (srv *Server) getFiles(ctx *gin.Context) {
	// TODO: Implement middleware
	files, err := srv.store.GetFiles(ctx, 1)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errJSON(err))
		return
	}

	ctx.JSON(http.StatusOK, files)
}
