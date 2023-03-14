package workers

import (
	"context"

	db "github.com/ShadrackAdwera/go-etl/db/sqlc"
	"github.com/hibiken/asynq"
)

type TaskProcessor interface {
	Start() error
	ProcessSendFileDataToDb(
		ctx context.Context,
		task *asynq.Task,
	) error
}

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

func (p *RedisTaskProcessor) Start() error {
	mux := asynq.NewServeMux()

	mux.HandleFunc(TaskSendFileDataToDb, p.ProcessSendFileDataToDb)

	return p.server.Run(mux)
}
