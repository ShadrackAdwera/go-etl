package db

import (
	"database/sql"
	"fmt"
	"os"
	"testing"

	"github.com/ShadrackAdwera/go-etl/utils"
	_ "github.com/lib/pq"
)

var testDb *sql.DB
var testQuery *Queries

func TestMain(m *testing.M) {
	var err error
	config, err := utils.LoadConfig("../..")

	if err != nil {
		panic(fmt.Errorf("error reading config: %v", err))
	}

	testDb, err = sql.Open(config.DbDriver, config.DbUrl)
	testQuery = New(testDb)
	if err != nil {
		panic(fmt.Errorf("error connecting to PG: %v", err))
	}
	os.Exit(m.Run())
}
