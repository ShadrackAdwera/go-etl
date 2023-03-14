package workers

import "github.com/hibiken/asynq"

type TaskDistributor interface{}

type RedisTaskDistributor struct {
	client *asynq.Client
}

func NewRedisTaskDistributor(rOpts asynq.RedisClientOpt) TaskDistributor {
	client := asynq.NewClient(rOpts)

	return &RedisTaskDistributor{
		client: client,
	}
}
