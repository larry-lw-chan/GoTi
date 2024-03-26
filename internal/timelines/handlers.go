package timelines

import (
	"net/http"

	"github.com/larry-lw-chan/goti/database"
	"github.com/larry-lw-chan/goti/internal/threads"
	"github.com/larry-lw-chan/goti/internal/utils/render"
)

func TimelineIndexHandler(w http.ResponseWriter, r *http.Request) {
	data := r.Context().Value("data").(map[string]any)

	// Algo - Get all Threads
	queries := threads.New(database.DB)
	threads, err := queries.GetAllThreads(r.Context())

	if err != nil {
		data["Threads"] = nil
	} else {
		data["Threads"] = threads
	}

	render.Template(w, data, "/timelines/index.app.tmpl", "/threads/partials/thread.app.tmpl")
}
