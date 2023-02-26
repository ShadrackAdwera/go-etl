-- name: CreateFile :one
INSERT INTO files ( 
    file_url, 
    created_by_id
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: GetFiles :many
SELECT * 
FROM files
WHERE created_by_id = $1;