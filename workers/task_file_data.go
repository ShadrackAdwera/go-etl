package workers

import (
	"context"
	"encoding/json"
	"fmt"

	db "github.com/ShadrackAdwera/go-etl/db/sqlc"
	"github.com/hibiken/asynq"
	"github.com/rs/zerolog/log"
)

type DistroSendFileToDbPayload struct {
	Matches []db.CreateMatchDataParams `json:"matches"`
}

const TaskSendFileDataToDb = "task:send_file_data_to_DB"

func (distro *RedisTaskDistributor) DistroSendFileDataToDb(
	ctx context.Context,
	payload *DistroSendFileToDbPayload,
	options ...asynq.Option,
) error {

	dt, err := json.Marshal(payload)

	if err != nil {
		return fmt.Errorf("unable to marshall json payload : %w", err)
	}

	task := asynq.NewTask(TaskSendFileDataToDb, dt, options...)

	info, err := distro.client.EnqueueContext(ctx, task)

	if err != nil {
		return fmt.Errorf("unable to enqueue task context : %w", err)
	}

	log.Info().Str("type", task.Type()).Str("queue_name", info.Queue).Msg("enqueued task")

	return nil
}

func (p *RedisTaskProcessor) ProcessSendFileDataToDb(
	ctx context.Context,
	task *asynq.Task,
) error {
	var matchData []db.CreateMatchDataParams

	if err := json.Unmarshal(task.Payload(), &matchData); err != nil {
		return fmt.Errorf("unable to unmarshall json body: %w", err)
	}

	for _, match := range matchData {
		_, err := p.store.CreateMatchData(ctx, db.CreateMatchDataParams{
			HomeScored:  match.HomeScored,
			AwayScored:  match.AwayScored,
			HomeTeam:    match.HomeTeam,
			AwayTeam:    match.AwayTeam,
			MatchDate:   match.MatchDate,
			Referee:     match.Referee,
			Winner:      match.Winner,
			Season:      match.Season,
			CreatedByID: match.CreatedByID,
			FileID:      match.FileID,
		})

		if err != nil {
			return fmt.Errorf("unable to add match to db : %v", match)
		}
	}

	log.Info().Msg("all tasks have been processed . . . ")

	return nil
}
