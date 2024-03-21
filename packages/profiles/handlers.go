package profiles

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"time"

	"github.com/larry-lw-chan/goti/database"
	"github.com/larry-lw-chan/goti/packages/sessions/flash"
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
			Private:   sql.NullInt64{Int64: 0, Valid: true},
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
	render.Template(w, data, "/profiles/show.app.tmpl")
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
	render.Template(w, data, "/profiles/edit.app.tmpl")
}

func EditPostHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	// Get user session information
	userSession := users.GetUserSession(r)
	queries := New(database.DB)

	// Check if user is private
	var private int64 = 0
	if r.FormValue("private") == "private" {
		private = 1
	}

	// Set Profile parameters
	updateProfileParams := UpdateProfileParams{
		Name:      sql.NullString{String: r.FormValue("name"), Valid: true},
		Bio:       sql.NullString{String: r.FormValue("bio"), Valid: true},
		Link:      sql.NullString{String: r.FormValue("link"), Valid: true},
		Private:   sql.NullInt64{Int64: private, Valid: true},
		UpdatedAt: time.Now().String(),
		UserID:    userSession.Id,
	}

	// Update profile
	_, err := queries.UpdateProfile(context.Background(), updateProfileParams)
	if err != nil {
		log.Println(err)
		flash.Set(w, r, flash.ERROR, "Profile update failed")
	} else {
		flash.Set(w, r, flash.SUCCESS, "Profile successfully updated")
	}

	// Do more stuff later
	http.Redirect(w, r, "/profiles/show", http.StatusSeeOther)
}
