package timelines

import (
	"net/http"

	"github.com/larry-lw-chan/goti/database"
	"github.com/larry-lw-chan/goti/internal/sessions/flash"
	"github.com/larry-lw-chan/goti/internal/threads"
	"github.com/larry-lw-chan/goti/internal/utils/render"
)

func TimelineIndexHandler(w http.ResponseWriter, r *http.Request) {
	data := r.Context().Value("data").(map[string]any)

	// Temp Solution - Get all Threads
	queries := threads.New(database.DB)
	threads, err := queries.GetAllThreads(r.Context())

	if err != nil {
		// Handle Error
		flash.Set(w, r, flash.ERROR, "Failed to get threads.  Please contact support.")
	} else {
		data["Threads"] = threads
	}

	render.Template(w, data, "/timelines/index.app.tmpl", "/partials/thread.app.tmpl")
}
