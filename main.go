package main

import (
	"database/sql"
	"fmt"
	"os"

	api "github.com/ShadrackAdwera/go-etl/api"
	db "github.com/ShadrackAdwera/go-etl/db/sqlc"
	"github.com/ShadrackAdwera/go-etl/utils"
	"github.com/ShadrackAdwera/go-etl/workers"
	"github.com/hibiken/asynq"
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

	fmt.Printf("Redis Address : %s\n", config.RedisAddress)

	store := db.NewStore(conn)

	redisOpts := asynq.RedisClientOpt{
		Addr: config.RedisAddress,
	}

	distro := workers.NewRedisTaskDistributor(redisOpts)

	srv := api.NewServer(store, config, distro)

	go runTaskProcessor(redisOpts, store)

	err = srv.StartServer(config.ServerAddress)

	if err != nil {
		log.Fatal().Err(err).Msg(fmt.Sprintf("unable to start server on address %s", config.ServerAddress))
		return
	}

}

func runTaskProcessor(opts asynq.RedisClientOpt, store db.TxStore) {
	processor := workers.NewRedisTaskProcessor(opts, store)

	err := processor.Start()

	if err != nil {
		log.Fatal().Err(err).Msg("unable to start task processor")
	}

	log.Info().Msg("started task processor")
}
