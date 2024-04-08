package threads

import (
	"context"
	"time"

	"github.com/larry-lw-chan/goti/database"
)

/****************************************************************
* Data Handling Logic
****************************************************************/
// Algo - Like or Unlike a thread
func likeOrUnlikeThread(context context.Context, threadId, userId int64) (message string, err error) {
	queries := New(database.DB)

	// Get all likes for a thread
	likes, err := queries.GetLikes(context, GetLikesParams{
		UserID:   userId,
		ThreadID: threadId,
	})
	if err != nil {
		return "", err
	}

	// If user has not liked the thread, then like the thread
	if len(likes) == 0 {
		queries.InsertLike(context, InsertLikeParams{
			ThreadID:  threadId,
			UserID:    userId,
			CreatedAt: time.Now().String(),
			UpdatedAt: time.Now().String(),
		})
		return "Liked thread", nil
	} else {
		queries.DeleteLike(context, DeleteLikeParams{
			ThreadID: threadId,
			UserID:   userId,
		})
		return "Deleted Like", nil
	}
}
