package api

import (
	"fmt"

	db "github.com/ShadrackAdwera/go-etl/db/sqlc"
	"github.com/ShadrackAdwera/go-etl/token"
	"github.com/ShadrackAdwera/go-etl/utils"
	"github.com/ShadrackAdwera/go-etl/workers"
	"github.com/gin-gonic/gin"
)

type Server struct {
	store      db.TxStore
	router     *gin.Engine
	tokenMaker token.TokenMaker
	taskDistro workers.TaskDistributor
}

func NewServer(store db.TxStore, config utils.AppConfig, distro workers.TaskDistributor) *Server {

	tokenMaker, err := token.NewPasetoMaker(config.PasetoKey)

	if err != nil {
		panic(err)
	}

	server := &Server{
		store:      store,
		tokenMaker: tokenMaker,
		taskDistro: distro,
	}

	router := gin.Default()

	router.POST("/api/sign-up", server.signUp)
	router.POST("/api/login", server.login)

	// auth middleware
	authRoutes := router.Group("/").Use(authMiddleware(tokenMaker))
	authRoutes.GET("/api/files", server.getFiles)
	authRoutes.POST("/api/files", server.uploadCsvFile)
	authRoutes.GET("/api/matches", server.getMatches)

	server.router = router
	return server
}

func (s *Server) StartServer(serverAddress string) error {
	return s.router.Run(serverAddress)
}

func errJSON(err error) gin.H {
	return gin.H{"message": fmt.Sprintf("error occured: %v", err)}
}
