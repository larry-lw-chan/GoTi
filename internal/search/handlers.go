package search

import (
	"net/http"

	"github.com/larry-lw-chan/goti/database"
	"github.com/larry-lw-chan/goti/internal/threads"
	"github.com/larry-lw-chan/goti/internal/utils/render"
)

func IndexThreadHandler(w http.ResponseWriter, r *http.Request) {
	data := r.Context().Value("data").(map[string]interface{})

	// Get all threads -- temporary
	queries := threads.New(database.DB)
	threads, err := queries.GetAllThreads(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Load data to template
	data["Threads"] = threads[:5]
	render.Template(w, data, "/search/index.app.tmpl", "/threads/__thread.app.tmpl")
}
