package profiles

import (
	"context"
	"database/sql"
	"net/http"
	"time"

	"github.com/larry-lw-chan/goti/database"
	"github.com/larry-lw-chan/goti/packages/users"
	"github.com/larry-lw-chan/goti/packages/utils/render"
)

func ProfileHandler(w http.ResponseWriter, r *http.Request) {
	data := r.Context().Value("data").(map[string]interface{})

	// Get user session information
	userSession := users.GetUserSession(r)
	queries := New(database.DB)

	// Get profile information or create if not exist
	profile, err := queries.GetProfileFromUserId(context.Background(), userSession.Id)
	if err != nil {
		profileParem := CreateProfileParams{
			Name:      sql.NullString{String: "please update your name", Valid: true},
			Bio:       sql.NullString{String: "please add a bio", Valid: true},
			Link:      sql.NullString{String: "please add your site link", Valid: true},
			UserID:    userSession.Id,
			CreatedAt: time.Now().String(),
			UpdatedAt: time.Now().String(),
		}
		profile, _ = queries.CreateProfile(context.Background(), profileParem)
	}
	data["Username"] = userSession.Username
	data["Profile"] = profile

	// Load user into profile
	render.Template(w, "/profiles/profile.app.tmpl", data)
}
