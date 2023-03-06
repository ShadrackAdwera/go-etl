package main

import (
	"database/sql"
	"fmt"
	"os"

	api "github.com/ShadrackAdwera/go-etl/api"
	db "github.com/ShadrackAdwera/go-etl/db/sqlc"
	"github.com/ShadrackAdwera/go-etl/utils"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	_ "github.com/lib/pq"
)

func main() {
	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatal().Err(err).Msg("unable to load config file")
		return
	}
	if config.Environment == "dev" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}

	conn, err := sql.Open(config.DbDriver, config.DbUrl)

	if err != nil {
		log.Fatal().Err(err).Msg("unable to connect to pg")
		return
	}

	store := db.NewStore(conn)
	srv := api.NewServer(store, config)
	err = srv.StartServer(config.ServerAddress)

	if err != nil {
		log.Fatal().Err(err).Msg(fmt.Sprintf("unable to start server on address %s", config.ServerAddress))
		return
	}

}
