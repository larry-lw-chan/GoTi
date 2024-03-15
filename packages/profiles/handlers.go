package profiles

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"time"

	"github.com/larry-lw-chan/goti/database"
	"github.com/larry-lw-chan/goti/packages/users"
	"github.com/larry-lw-chan/goti/packages/utils/render"
)

func ShowHandler(w http.ResponseWriter, r *http.Request) {
	data := r.Context().Value("data").(map[string]interface{})

	// Get user session information
	userSession := users.GetUserSession(r)
	queries := New(database.DB)

	// Get profile information or create if not exist
	profile, err := queries.GetProfileFromUserId(context.Background(), userSession.Id)
	if err != nil {
		profileParem := CreateProfileParams{
			Name:      sql.NullString{String: "your name", Valid: true},
			Bio:       sql.NullString{String: "add bio", Valid: true},
			Link:      sql.NullString{String: "add links", Valid: true},
			Avatar:    sql.NullString{Valid: false},
			UserID:    userSession.Id,
			CreatedAt: time.Now().String(),
			UpdatedAt: time.Now().String(),
		}
		profile, _ = queries.CreateProfile(context.Background(), profileParem)
	}
	data["Username"] = userSession.Username
	data["Profile"] = profile

	// Load user into profile
	render.Template(w, "/profiles/show.app.tmpl", data)
}

func EditHandler(w http.ResponseWriter, r *http.Request) {
	data := r.Context().Value("data").(map[string]interface{})

	// Get user session information
	userSession := users.GetUserSession(r)
	queries := New(database.DB)

	// Get profile information
	profile, err := queries.GetProfileFromUserId(context.Background(), userSession.Id)
	if err != nil {
		log.Println(err)
	}

	// Load username and profile to pass to template
	data["Username"] = userSession.Username
	data["Profile"] = profile

	// Load user into profile
	render.Template(w, "/profiles/edit.app.tmpl", data)
}

func EditPostHandler(w http.ResponseWriter, r *http.Request) {
	// Do more stuff later
	http.Redirect(w, r, "/profiles/show", http.StatusSeeOther)
}
