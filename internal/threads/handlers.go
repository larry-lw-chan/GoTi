package threads

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/larry-lw-chan/goti/database"
	"github.com/larry-lw-chan/goti/internal/auth"
	"github.com/larry-lw-chan/goti/internal/helper"
	"github.com/larry-lw-chan/goti/internal/sessions/flash"
	"github.com/larry-lw-chan/goti/internal/utils/render"
)

/********************************************************
* Threads
*********************************************************/
func IndexThreadHandler(w http.ResponseWriter, r *http.Request) {
	data := r.Context().Value("data").(map[string]any)

	// Algo - Get all Threads
	queries := New(database.DB)
	threads, err := queries.GetAllThreads(r.Context())

	if err != nil {
		data["Threads"] = nil
	} else {
		data["Threads"] = threads
	}

	render.Template(w, data,
		"/threads/index.app.tmpl",
		"/threads/__avatar.app.tmpl",
		"/threads/__username.app.tmpl",
		"/threads/__thread_content.app.tmpl",
		"/threads/__menu_icon_status.app.tmpl")
}

func NewThreadHandler(w http.ResponseWriter, r *http.Request) {
	data := r.Context().Value("data").(map[string]interface{})
	render.Template(w, data, "/threads/new.app.tmpl")
}

func NewPostThreadHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	// Handle Form Validation
	errs := validateNewThread(r)
	if errs != nil {
		messages := helper.GetErrorMessages(errs)
		flash.Set(w, r, flash.ERROR, messages)
		http.Redirect(w, r, "/threads/new", http.StatusSeeOther)
		return
	}

	// Get user session information
	userSession := auth.GetUserSession(r)

	// Create new thread
	queries := New(database.DB)

	threadParam := CreateThreadParams{
		Content:   r.FormValue("content"),
		ThreadID:  sql.NullInt64{Valid: false},
		UserID:    userSession.ID,
		CreatedAt: time.Now().String(),
		UpdatedAt: time.Now().String(),
	}

	thread, err := queries.CreateThread(r.Context(), threadParam)

	if err != nil {
		flash.Set(w, r, flash.ERROR, "Failed to create new thread.  Please contact support.")
		http.Redirect(w, r, "/threads/new", http.StatusSeeOther)
		return
	}

	// Set flash message and redirect to thread page
	thread_path := "/threads/" + fmt.Sprint(thread.ID)
	flash.Set(w, r, flash.SUCCESS, "New thread created.")
	http.Redirect(w, r, thread_path, http.StatusSeeOther)
}

func ShowThreadHandler(w http.ResponseWriter, r *http.Request) {
	data := r.Context().Value("data").(map[string]interface{})

	// Get thread_id from URL
	thread_id := chi.URLParam(r, "thread_id")
	id, err := strconv.ParseInt(thread_id, 10, 64)
	if err != nil {
		// Handle Error
		flash.Set(w, r, flash.ERROR, "Failed to get thread.  Please contact support.")
		http.Redirect(w, r, "/threads", http.StatusSeeOther)
		return
	}

	// Get thread from id
	queries := New(database.DB)
	thread, err := queries.GetThreadByID(r.Context(), int64(id))
	if err != nil {
		// Handle Error
		flash.Set(w, r, flash.ERROR, "Failed to get thread.  Please contact support.")
		http.Redirect(w, r, "/threads", http.StatusSeeOther)
		return
	}

	// Move stuff to thread
	data["Thread"] = thread

	render.Template(w, data,
		"/threads/show.app.tmpl",
		"/threads/__avatar.app.tmpl",
		"/threads/__username.app.tmpl",
		"/threads/__thread_content.app.tmpl",
		"/threads/__menu_icon_status.app.tmpl",
	)
}

/********************************************************
* Likes
*********************************************************/
func LikeThreadHandler(w http.ResponseWriter, r *http.Request) {
	data := r.Context().Value("data").(map[string]interface{})

	// Get userID from sessions
	userSession := auth.GetUserSession(r)
	userId := userSession.ID

	// Get thread_id from URL
	threadIdStr := chi.URLParam(r, "thread_id")
	threadId, err := strconv.ParseInt(threadIdStr, 10, 64)
	handleLikeError(w, r, err)

	// Check if user already liked thread
	likeStatus, err := likeOrUnlikeThread(r.Context(), threadId, userId)
	handleLikeError(w, r, err)

	// Get Likes Count
	likeCount, err := getLikeCounts(r.Context(), threadId)
	handleLikeError(w, r, err)

	// Render Result
	data["LikeStatus"] = likeStatus
	data["LikeCount"] = likeCount
	data["ThreadId"] = threadId
	render.Partial(w, data, "/threads/thread_icon_status.app.tmpl")
}

func handleLikeError(w http.ResponseWriter, r *http.Request, err error) {
	if err != nil {
		// Handle Error
		flash.Set(w, r, flash.ERROR, "Failed to like thread.  Please contact support.")
		http.Redirect(w, r, "/threads", http.StatusSeeOther)
		return
	}
}

/********************************************************
* Partials
*********************************************************/
func UserThreadsHandler(w http.ResponseWriter, r *http.Request) {
	data := r.Context().Value("data").(map[string]interface{})

	// Get user_id from URL
	userId := chi.URLParam(r, "user_id")
	userID, err := strconv.ParseInt(userId, 10, 64)
	if err != nil {
		// Handle Error
		flash.Set(w, r, flash.ERROR, "Failed to get threads.  Please contact support.")
		data["Threads"] = nil
	}

	// Temp Solution - Get all Threads
	queries := New(database.DB)
	threads, err := queries.GetUserThreads(r.Context(), userID)

	if err != nil {
		// Handle Error
		// flash.Set(w, r, flash.ERROR, "Failed to get threads.  Please contact support.")
		data["Threads"] = nil
	} else {
		data["Threads"] = threads
	}

	render.Partial(w, data,
		"/threads/user_threads.app.tmpl",
		"/threads/__avatar.app.tmpl",
		"/threads/__username.app.tmpl",
		"/threads/__thread_content.app.tmpl",
		"/threads/__menu_icon_status.app.tmpl")
}

func UserRepliesHandler(w http.ResponseWriter, r *http.Request) {
	data := r.Context().Value("data").(map[string]interface{})

	// Get user_id from URL
	userId := chi.URLParam(r, "user_id")
	userID, err := strconv.ParseInt(userId, 10, 64)
	if err != nil {
		// Handle Error
		flash.Set(w, r, flash.ERROR, "Failed to get threads.  Please contact support.")
		data["Threads"] = nil
	}

	// Temp Solution - Get all Threads
	queries := New(database.DB)
	threads, err := queries.GetUserThreads(r.Context(), userID)

	if err != nil {
		// Handle Error
		// flash.Set(w, r, flash.ERROR, "Failed to get threads.  Please contact support.")
		data["Threads"] = nil
	} else {
		data["Threads"] = threads
	}
	render.Partial(w, data,
		"/threads/user_replies.app.tmpl",
		"/threads/__avatar.app.tmpl",
		"/threads/__username.app.tmpl",
		"/threads/__thread_content.app.tmpl",
		"/threads/__menu_icon_status.app.tmpl")
}

func UserRepostHandler(w http.ResponseWriter, r *http.Request) {
	data := r.Context().Value("data").(map[string]interface{})

	// Get user_id from URL
	userId := chi.URLParam(r, "user_id")
	userID, err := strconv.ParseInt(userId, 10, 64)
	if err != nil {
		// Handle Error
		flash.Set(w, r, flash.ERROR, "Failed to get threads.  Please contact support.")
		data["Threads"] = nil
	}

	// Temp Solution - Get all Threads
	queries := New(database.DB)
	threads, err := queries.GetUserThreads(r.Context(), userID)

	if err != nil {
		data["Threads"] = nil
	} else {
		data["Threads"] = threads
	}

	render.Partial(w, data,
		"/threads/user_repost.app.tmpl",
		"/threads/__avatar.app.tmpl",
		"/threads/__username.app.tmpl",
		"/threads/__thread_content.app.tmpl",
		"/threads/__menu_icon_status.app.tmpl")
}
