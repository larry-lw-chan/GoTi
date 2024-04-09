package threads

import (
	"context"
	"database/sql"
	"time"

	"github.com/larry-lw-chan/goti/database"
)

/****************************************************************
* Data Handling Logic
****************************************************************/
// Check to see if user has liked a thread
func didUserLikedThread(context context.Context, threadId, userId int64) (likeStatus bool) {
	queries := New(database.DB)
	_, err := queries.CheckIfUserLikedThread(context, CheckIfUserLikedThreadParams{
		UserID:   userId,
		ThreadID: threadId,
	})
	if err != sql.ErrNoRows {
		return true
	} else {
		return false
	}
}

// Like or Unlike a thread
func likeOrUnlikeThread(context context.Context, threadId, userId int64) (likeStatus bool, err error) {
	queries := New(database.DB)

	// See if user has liked the thread
	liked := didUserLikedThread(context, threadId, userId)

	if !liked {
		// If user has not liked the thread, then like the thread
		queries.InsertLike(context, InsertLikeParams{
			ThreadID:  threadId,
			UserID:    userId,
			CreatedAt: time.Now().String(),
			UpdatedAt: time.Now().String(),
		})
		likeStatus = true
	} else {
		// If user has liked the thread, then unlike the thread
		queries.DeleteLike(context, DeleteLikeParams{
			ThreadID: threadId,
			UserID:   userId,
		})
		likeStatus = false
	}

	return likeStatus, nil
}

func getLikeCounts(context context.Context, threadId int64) (likeCount int64, err error) {
	queries := New(database.DB)
	likeCount, err = queries.GetThreadLikeCount(context, threadId)
	if err != nil {
		return 0, err
	}
	return likeCount, nil
}
