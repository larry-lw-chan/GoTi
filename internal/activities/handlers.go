package activities

import (
	"net/http"

	"github.com/larry-lw-chan/goti/database"
	"github.com/larry-lw-chan/goti/internal/profiles"
	"github.com/larry-lw-chan/goti/internal/utils/render"
)

func IndexActivityHandler(w http.ResponseWriter, r *http.Request) {
	render.Template(w, nil, "/activities/index.app.tmpl")
}

/****************************************************************
* Profile Partials
****************************************************************/
func GetAllHandler(w http.ResponseWriter, r *http.Request) {
	data := r.Context().Value("data").(map[string]interface{})

	// Get all profiles
	queries := profiles.New(database.DB)
	profiles, err := queries.GetAllProfiles(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Load profiles to pass to template
	data["Profiles"] = profiles
	render.Partial(w, data, "/activities/all.app.tmpl", "/activities/__profile.app.tmpl")
}

func GetFollowHandler(w http.ResponseWriter, r *http.Request) {
	data := r.Context().Value("data").(map[string]interface{})

	// Get all profiles
	queries := profiles.New(database.DB)
	profiles, err := queries.GetAllProfiles(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Load profiles to pass to template
	data["Profiles"] = profiles
	render.Partial(w, data, "/activities/all.app.tmpl", "/activities/__profile.app.tmpl")
}

func GetRepliesHandler(w http.ResponseWriter, r *http.Request) {
	data := r.Context().Value("data").(map[string]interface{})

	// Get all profiles
	queries := profiles.New(database.DB)
	profiles, err := queries.GetAllProfiles(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Load profiles to pass to template
	data["Profiles"] = profiles
	render.Partial(w, data, "/activities/all.app.tmpl", "/activities/__profile.app.tmpl")
}

func GetRepostHandler(w http.ResponseWriter, r *http.Request) {
	data := r.Context().Value("data").(map[string]interface{})

	// Get all profiles
	queries := profiles.New(database.DB)
	profiles, err := queries.GetAllProfiles(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Load profiles to pass to template
	data["Profiles"] = profiles
	render.Partial(w, data, "/activities/all.app.tmpl", "/activities/__profile.app.tmpl")
}
