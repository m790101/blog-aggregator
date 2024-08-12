-- name: CreateFeed :one
INSERT INTO feeds (id ,name, url, created_at, updated_at, user_id,last_fetch_at)
VALUES ($1,$2,$3,$4,$5,$6,$7)
RETURNING *;




-- name: GetFeed :many
SELECT * FROM feeds;



-- name: GetNextFeedToFetch :many
SELECT * FROM feeds
ORDER BY last_fetch_at ASC NULLS FIRST
LIMIT $1;


-- name: MarkFeedFetch :one
UPDATE feeds
SET last_fetch_at = NOW(),
updated_at = NOW()
WHERE id = $1
RETURNING *;





