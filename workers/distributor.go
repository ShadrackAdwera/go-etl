package workers

import (
	"context"

	"github.com/hibiken/asynq"
)

type TaskDistributor interface {
	DistroSendFileDataToDb(
		ctx context.Context,
		payload *DistroSendFileToDbPayload,
		options ...asynq.Option,
	) error
}

type RedisTaskDistributor struct {
	client *asynq.Client
}

func NewRedisTaskDistributor(rOpts asynq.RedisClientOpt) TaskDistributor {
	client := asynq.NewClient(rOpts)

	return &RedisTaskDistributor{
		client: client,
	}
}
