-- name: CreateMatchData :one
INSERT INTO match_data (
    home_scored, 
    away_scored, 
    home_team, 
    away_team,
    match_date,
    referee,
    winner,
    season,
    created_by_id,
    file_id
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8, $9, $10
)
RETURNING *;

-- name: GetMatchData :many
SELECT * 
FROM match_data
WHERE created_by_id = $1;

-- name: GetMatchDataById :one
SELECT * 
FROM match_data
WHERE id = $1;

-- name: UpdateMatchData :one
UPDATE match_data 
SET
  home_scored = COALESCE(sqlc.narg(home_scored),home_scored),
  away_scored = COALESCE(sqlc.narg(away_scored),away_scored),
  home_team = COALESCE(sqlc.narg(home_team),home_team),
  away_team = COALESCE(sqlc.narg(away_team),away_team),
  match_date = COALESCE(sqlc.narg(match_date),match_date),
  referee = COALESCE(sqlc.narg(referee),referee),
  winner = COALESCE(sqlc.narg(winner),winner),
  season = COALESCE(sqlc.narg(season),season),
  file_id = COALESCE(sqlc.narg(file_id),file_id)
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: DeleteMatchData :exec
DELETE FROM 
match_data 
WHERE id = $1;
