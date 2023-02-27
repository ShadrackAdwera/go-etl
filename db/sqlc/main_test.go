package db

import (
	"database/sql"
	"fmt"
	"os"
	"testing"

	"github.com/ShadrackAdwera/go-etl/token"
	"github.com/ShadrackAdwera/go-etl/utils"
	_ "github.com/lib/pq"
)

var testDb *sql.DB
var testQuery *Queries
var testMaker token.TokenMaker

func TestMain(m *testing.M) {
	var err error
	config, err := utils.LoadConfig("../..")

	if err != nil {
		panic(fmt.Errorf("error reading config: %v", err))
	}

	testMaker, err = token.NewPasetoMaker(config.PasetoKey)
	if err != nil {
		panic(fmt.Errorf("fail to create paseto token maker: %v", err))
	}

	testDb, err = sql.Open(config.DbDriver, config.DbUrl)
	testQuery = New(testDb)
	if err != nil {
		panic(fmt.Errorf("error connecting to PG: %v", err))
	}
	os.Exit(m.Run())
}
