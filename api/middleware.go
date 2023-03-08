package api

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/ShadrackAdwera/go-etl/token"
	"github.com/gin-gonic/gin"
)

const (
	authorization = "authorization"
	authType      = "bearer"
	authPayload   = "auth_payload"
)

func authMiddleware(tokenMaker token.TokenMaker) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader(authorization)

		if len(authHeader) == 0 {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errJSON(errors.New("provide the auth header")))
			return
		}

		fields := strings.Fields(authHeader)

		if strings.ToLower(fields[0]) != authType {
			err := fmt.Sprintf("unsupported auth type %s", fields[0])
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errJSON(errors.New(err)))
			return
		}

		payload, err := tokenMaker.VerifyToken(fields[1])

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errJSON(err))
			return
		}

		ctx.Set(authPayload, payload)
		ctx.Next()

	}
}
