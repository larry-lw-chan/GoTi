// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package threads

import (
	"database/sql"
)

type Hashtag struct {
	ID        int64
	Hashtag   string
	CreatedAt string
	UpdatedAt string
}

type Like struct {
	ID        int64
	ThreadID  int64
	UserID    int64
	CreatedAt string
	UpdatedAt string
	Foreign   interface{}
}

type Profile struct {
	ID        int64
	Username  string
	Name      sql.NullString
	Bio       sql.NullString
	Link      sql.NullString
	Avatar    sql.NullString
	Private   sql.NullInt64
	UserID    int64
	CreatedAt string
	UpdatedAt string
}

type Thread struct {
	ID        int64
	Content   string
	ThreadID  sql.NullInt64
	UserID    int64
	CreatedAt string
	UpdatedAt string
	Foreign   interface{}
}

type ThreadsHashtag struct {
	ID        int64
	ThreadID  sql.NullInt64
	HashtagID int64
	CreatedAt string
	UpdatedAt string
}
