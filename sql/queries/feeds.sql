-- name: CreateFeed :one
INSERT INTO feeds (
    id,
    created_at,
    updated_at,
    name,
    url,
    user_id
) VALUES (
    $1,
    $2,
    $3,
    $4,
    $5, 
    $6
)
RETURNING *;

-- name: GetFeed :one
SELECT * FROM feeds WHERE url = $1;

-- name: GetFeeds :many
SELECT * FROM feeds;

-- name: DeleteFeeds :exec
DELETE FROM feeds;

-- name: MarkFeedFetched :exec
UPDATE feeds SET last_fetched_at = $1, updated_at = $2
WHERE id = $3;

-- name: GetNextFeedToFetch :one
SELECT * FROM feeds
ORDER BY
    last_fetched_at NULLS FIRST,
    last_fetched_at ASC
LIMIT 1;