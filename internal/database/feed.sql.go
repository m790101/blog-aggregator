// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: feed.sql

package database

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const createFeed = `-- name: CreateFeed :one
INSERT INTO feeds (id ,name, url, created_at, updated_at, user_id)
VALUES ($1,$2,$3,$4,$5,$6)
RETURNING id, name, url, created_at, updated_at, user_id, last_fetch_at
`

type CreateFeedParams struct {
	ID        uuid.UUID
	Name      string
	Url       string
	CreatedAt time.Time
	UpdatedAt time.Time
	UserID    uuid.UUID
}

func (q *Queries) CreateFeed(ctx context.Context, arg CreateFeedParams) (Feed, error) {
	row := q.db.QueryRowContext(ctx, createFeed,
		arg.ID,
		arg.Name,
		arg.Url,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.UserID,
	)
	var i Feed
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Url,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.UserID,
		&i.LastFetchAt,
	)
	return i, err
}

const getFeed = `-- name: GetFeed :one
SELECT id, name, url, created_at, updated_at, user_id, last_fetch_at FROM feeds
`

func (q *Queries) GetFeed(ctx context.Context) (Feed, error) {
	row := q.db.QueryRowContext(ctx, getFeed)
	var i Feed
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Url,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.UserID,
		&i.LastFetchAt,
	)
	return i, err
}

const getNextFeedToFetch = `-- name: GetNextFeedToFetch :one
SELECT id, name, url, created_at, updated_at, user_id, last_fetch_at FROM feeds
ORDER BY last_fetch_at ASC NULLS FIRST
LIMIT 1
`

func (q *Queries) GetNextFeedToFetch(ctx context.Context) (Feed, error) {
	row := q.db.QueryRowContext(ctx, getNextFeedToFetch)
	var i Feed
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Url,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.UserID,
		&i.LastFetchAt,
	)
	return i, err
}

const markFeedFetch = `-- name: MarkFeedFetch :one
UPDATE feeds
SET last_fetch_at = NOW(),
updated_at = NOW()
WHERE id = $1
RETURNING id, name, url, created_at, updated_at, user_id, last_fetch_at
`

func (q *Queries) MarkFeedFetch(ctx context.Context, id uuid.UUID) (Feed, error) {
	row := q.db.QueryRowContext(ctx, markFeedFetch, id)
	var i Feed
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Url,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.UserID,
		&i.LastFetchAt,
	)
	return i, err
}
