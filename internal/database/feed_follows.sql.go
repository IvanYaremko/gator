// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: feed_follows.sql

package database

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const createFeedFollow = `-- name: CreateFeedFollow :one
WITH inserted_feed_follow AS (
        INSERT INTO feed_follows (
        id,
        created_at,
        updated_at,
        user_id,
        feed_id
    ) VALUES (
        $1,
        $2,
        $3,
        $4,
        $5
    )
     RETURNING id, created_at, updated_at, user_id, feed_id
)
SELECT
    inserted_feed_follow.id, inserted_feed_follow.created_at, inserted_feed_follow.updated_at, inserted_feed_follow.user_id, inserted_feed_follow.feed_id,
    feeds.name AS feed_name,
    users.name AS user_name
FROM inserted_feed_follow
INNER JOIN feeds ON inserted_feed_follow.feed_id = feeds.id
INNER JOIN users ON inserted_feed_follow.user_id = users.id
`

type CreateFeedFollowParams struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	UserID    uuid.UUID
	FeedID    uuid.UUID
}

type CreateFeedFollowRow struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	UserID    uuid.UUID
	FeedID    uuid.UUID
	FeedName  string
	UserName  string
}

func (q *Queries) CreateFeedFollow(ctx context.Context, arg CreateFeedFollowParams) (CreateFeedFollowRow, error) {
	row := q.db.QueryRowContext(ctx, createFeedFollow,
		arg.ID,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.UserID,
		arg.FeedID,
	)
	var i CreateFeedFollowRow
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.UserID,
		&i.FeedID,
		&i.FeedName,
		&i.UserName,
	)
	return i, err
}

const deleteFeedFollows = `-- name: DeleteFeedFollows :exec
DELETE FROM feed_follows
`

func (q *Queries) DeleteFeedFollows(ctx context.Context) error {
	_, err := q.db.ExecContext(ctx, deleteFeedFollows)
	return err
}

const deleteFeedFollowsById = `-- name: DeleteFeedFollowsById :exec
DELETE from feed_follows
WHERE user_id = $1 AND feed_id = $2
`

type DeleteFeedFollowsByIdParams struct {
	UserID uuid.UUID
	FeedID uuid.UUID
}

func (q *Queries) DeleteFeedFollowsById(ctx context.Context, arg DeleteFeedFollowsByIdParams) error {
	_, err := q.db.ExecContext(ctx, deleteFeedFollowsById, arg.UserID, arg.FeedID)
	return err
}

const getFeedFollowsForUsers = `-- name: GetFeedFollowsForUsers :many
SELECT feed_follows.id, feed_follows.created_at, feed_follows.updated_at, feed_follows.user_id, feed_follows.feed_id, users.name AS user_name, feeds.name as feed_name
FROM feed_follows
INNER JOIN users ON feed_follows.user_id = users.id
INNER JOIN feeds on feed_follows.feed_id = feeds.id
WHERE feed_follows.user_id = $1
`

type GetFeedFollowsForUsersRow struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	UserID    uuid.UUID
	FeedID    uuid.UUID
	UserName  string
	FeedName  string
}

func (q *Queries) GetFeedFollowsForUsers(ctx context.Context, userID uuid.UUID) ([]GetFeedFollowsForUsersRow, error) {
	rows, err := q.db.QueryContext(ctx, getFeedFollowsForUsers, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetFeedFollowsForUsersRow
	for rows.Next() {
		var i GetFeedFollowsForUsersRow
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.UserID,
			&i.FeedID,
			&i.UserName,
			&i.FeedName,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}