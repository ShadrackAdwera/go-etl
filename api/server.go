package api

import (
	"fmt"

	db "github.com/ShadrackAdwera/go-etl/db/sqlc"
	"github.com/ShadrackAdwera/go-etl/token"
	"github.com/ShadrackAdwera/go-etl/utils"
	"github.com/gin-gonic/gin"
)

type Server struct {
	store      db.TxStore
	router     *gin.Engine
	config     utils.AppConfig
	tokenMaker token.TokenMaker
}

func NewServer(store db.TxStore, config utils.AppConfig) *Server {

	tokenMaker, err := token.NewPasetoMaker(config.PasetoKey)

	if err != nil {
		panic(err)
	}

	server := &Server{
		store:      store,
		tokenMaker: tokenMaker,
	}

	router := gin.Default()

	router.POST("/api/sign-up", server.signUp)
	router.POST("/api/login", server.login)

	server.router = router
	return server
}

func errJSON(err error) gin.H {
	return gin.H{"message": fmt.Sprintf("error occured: %v", err)}
}
