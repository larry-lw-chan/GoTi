// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: query.sql

package threads

import (
	"context"
	"database/sql"
)

const createThread = `-- name: CreateThread :one
INSERT INTO threads (content, thread_id, user_id, created_at, updated_at)
VALUES (?, ?, ?, ?, ?)
RETURNING id, content, thread_id, user_id, created_at, updated_at, "foreign"
`

type CreateThreadParams struct {
	Content   string
	ThreadID  sql.NullInt64
	UserID    int64
	CreatedAt string
	UpdatedAt string
}

func (q *Queries) CreateThread(ctx context.Context, arg CreateThreadParams) (Thread, error) {
	row := q.db.QueryRowContext(ctx, createThread,
		arg.Content,
		arg.ThreadID,
		arg.UserID,
		arg.CreatedAt,
		arg.UpdatedAt,
	)
	var i Thread
	err := row.Scan(
		&i.ID,
		&i.Content,
		&i.ThreadID,
		&i.UserID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Foreign,
	)
	return i, err
}

const deleteLike = `-- name: DeleteLike :one
DELETE FROM likes
WHERE thread_id = ? AND user_id = ?
RETURNING id, thread_id, user_id, created_at, updated_at, "foreign"
`

type DeleteLikeParams struct {
	ThreadID int64
	UserID   int64
}

func (q *Queries) DeleteLike(ctx context.Context, arg DeleteLikeParams) (Like, error) {
	row := q.db.QueryRowContext(ctx, deleteLike, arg.ThreadID, arg.UserID)
	var i Like
	err := row.Scan(
		&i.ID,
		&i.ThreadID,
		&i.UserID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Foreign,
	)
	return i, err
}

const getAllThreads = `-- name: GetAllThreads :many
SELECT threads.id, content, username, avatar, (
    SELECT COUNT(*) 
    FROM likes 
    WHERE likes.thread_id=threads.id
) AS likes
FROM threads
JOIN profiles ON profiles.user_id = threads.user_id
ORDER BY threads.created_at desc
`

type GetAllThreadsRow struct {
	ID       int64
	Content  string
	Username string
	Avatar   sql.NullString
	Likes    int64
}

func (q *Queries) GetAllThreads(ctx context.Context) ([]GetAllThreadsRow, error) {
	rows, err := q.db.QueryContext(ctx, getAllThreads)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetAllThreadsRow
	for rows.Next() {
		var i GetAllThreadsRow
		if err := rows.Scan(
			&i.ID,
			&i.Content,
			&i.Username,
			&i.Avatar,
			&i.Likes,
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

const getLikes = `-- name: GetLikes :many
SELECT id, thread_id, user_id, created_at, updated_at, "foreign"
FROM likes
WHERE thread_id = ? AND user_id = ?
`

type GetLikesParams struct {
	ThreadID int64
	UserID   int64
}

func (q *Queries) GetLikes(ctx context.Context, arg GetLikesParams) ([]Like, error) {
	rows, err := q.db.QueryContext(ctx, getLikes, arg.ThreadID, arg.UserID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Like
	for rows.Next() {
		var i Like
		if err := rows.Scan(
			&i.ID,
			&i.ThreadID,
			&i.UserID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Foreign,
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

const getThreadByID = `-- name: GetThreadByID :one
SELECT threads.id, content, username, avatar
FROM threads
JOIN profiles ON profiles.user_id = threads.user_id
WHERE threads.id = ?
`

type GetThreadByIDRow struct {
	ID       int64
	Content  string
	Username string
	Avatar   sql.NullString
}

func (q *Queries) GetThreadByID(ctx context.Context, id int64) (GetThreadByIDRow, error) {
	row := q.db.QueryRowContext(ctx, getThreadByID, id)
	var i GetThreadByIDRow
	err := row.Scan(
		&i.ID,
		&i.Content,
		&i.Username,
		&i.Avatar,
	)
	return i, err
}

const getUserThreads = `-- name: GetUserThreads :many
SELECT threads.id, content, username, avatar, (
    SELECT COUNT(*) 
    FROM likes 
    WHERE likes.thread_id=threads.id
) AS likes
FROM threads
JOIN profiles ON profiles.user_id = threads.user_id
WHERE threads.user_id = ?
ORDER BY threads.created_at desc
`

type GetUserThreadsRow struct {
	ID       int64
	Content  string
	Username string
	Avatar   sql.NullString
	Likes    int64
}

func (q *Queries) GetUserThreads(ctx context.Context, userID int64) ([]GetUserThreadsRow, error) {
	rows, err := q.db.QueryContext(ctx, getUserThreads, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetUserThreadsRow
	for rows.Next() {
		var i GetUserThreadsRow
		if err := rows.Scan(
			&i.ID,
			&i.Content,
			&i.Username,
			&i.Avatar,
			&i.Likes,
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

const insertLike = `-- name: InsertLike :one
INSERT INTO likes (thread_id, user_id, created_at, updated_at)
VALUES (?, ?, ?, ?)
RETURNING id, thread_id, user_id, created_at, updated_at, "foreign"
`

type InsertLikeParams struct {
	ThreadID  int64
	UserID    int64
	CreatedAt string
	UpdatedAt string
}

func (q *Queries) InsertLike(ctx context.Context, arg InsertLikeParams) (Like, error) {
	row := q.db.QueryRowContext(ctx, insertLike,
		arg.ThreadID,
		arg.UserID,
		arg.CreatedAt,
		arg.UpdatedAt,
	)
	var i Like
	err := row.Scan(
		&i.ID,
		&i.ThreadID,
		&i.UserID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Foreign,
	)
	return i, err
}
