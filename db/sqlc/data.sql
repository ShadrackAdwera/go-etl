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