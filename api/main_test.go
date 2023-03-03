package api

import (
	"os"
	"testing"
	"time"

	db "github.com/ShadrackAdwera/go-etl/db/sqlc"
	"github.com/ShadrackAdwera/go-etl/utils"
	"github.com/gin-gonic/gin"
)

func newTestServer(t *testing.T, store db.TxStore) *Server {
	config := utils.AppConfig{
		PasetoKey:           utils.RandomString(32),
		AccessTokenDuration: time.Minute,
	}

	server := NewServer(store, config)
	return server
}

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	os.Exit(m.Run())
}
