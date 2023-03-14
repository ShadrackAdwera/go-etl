package workers

import (
	db "github.com/ShadrackAdwera/go-etl/db/sqlc"
	"github.com/hibiken/asynq"
)

type TaskProcessor interface{}

type RedisTaskProcessor struct {
	server *asynq.Server
	store  db.TxStore
}

func NewRedisTaskProcessor(rOpts asynq.RedisClientOpt, store db.TxStore) TaskProcessor {
	server := asynq.NewServer(rOpts, asynq.Config{})
	return &RedisTaskProcessor{
		server: server,
		store:  store,
	}
}
