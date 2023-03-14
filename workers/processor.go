package workers

import (
	"context"

	db "github.com/ShadrackAdwera/go-etl/db/sqlc"
	"github.com/hibiken/asynq"
	"github.com/rs/zerolog/log"
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
	server := asynq.NewServer(rOpts, asynq.Config{
		ErrorHandler: asynq.ErrorHandlerFunc(func(ctx context.Context, task *asynq.Task, err error) {
			log.Error().Err(err).Str("task_type", task.Type()).Bytes("payload", task.Payload()).Msg("task processing failed")
		}),
	})
	return &RedisTaskProcessor{
		server: server,
		store:  store,
	}
}

func (p *RedisTaskProcessor) Start() error {
	mux := asynq.NewServeMux()

	mux.HandleFunc(TaskSendFileDataToDb, p.ProcessSendFileDataToDb)

	return p.server.Start(mux)
}
