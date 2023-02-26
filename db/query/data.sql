-- name: CreateData :one
INSERT INTO data (
    home_scored, 
    away_scored, 
    home_team, 
    away_team,
    match_day, 
    referee,
    is_blocked, 
    winner,
    season,
    created_by_id,
    file_id
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11
)
RETURNING *;

-- name: GetData :many
SELECT * 
FROM data
WHERE created_by_id = $1;

-- name: GetDataById :one
SELECT * 
FROM data
WHERE id = $1;

-- name: UpdateData :one
UPDATE data 
SET
  home_scored = COALESCE(sqlc.narg(home_scored),home_scored),
  away_scored = COALESCE(sqlc.narg(away_scored),away_scored),
  home_team = COALESCE(sqlc.narg(home_team),home_team),
  away_team = COALESCE(sqlc.narg(away_team),away_team)
  match_day = COALESCE(sqlc.narg(match_day),match_day),
  referee = COALESCE(sqlc.narg(referee),referee),
  is_blocked = COALESCE(sqlc.narg(is_blocked),is_blocked),
  winner = COALESCE(sqlc.narg(winner),winner)
  season = COALESCE(sqlc.narg(season),season),
  file_id = COALESCE(sqlc.narg(file_id),file_id)
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: DeleteData :exec
DELETE FROM 
data 
WHERE id = $1;
