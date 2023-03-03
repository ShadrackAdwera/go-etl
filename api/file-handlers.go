package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserArgs struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (srv *Server) getFiles(ctx *gin.Context) {
	// TODO: Implement middleware

	var userId UserArgs

	if err := ctx.ShouldBindUri(&userId); err != nil {
		ctx.JSON(http.StatusBadRequest, errJSON(err))
		return
	}

	files, err := srv.store.GetFiles(ctx, userId.ID)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errJSON(err))
		return
	}

	ctx.JSON(http.StatusOK, files)
}
