package profiles

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"time"

	"github.com/larry-lw-chan/goti/database"
	"github.com/larry-lw-chan/goti/packages/filestore"
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

	// Redirect to profile page
	http.Redirect(w, r, "/profiles/show", http.StatusSeeOther)
}

func EditPhotoHandler(w http.ResponseWriter, r *http.Request) {
	data := r.Context().Value("data").(map[string]interface{})
	render.Template(w, data, "/profiles/edit-photo.app.tmpl")
}

// Todo: Implement photo upload
func EditPhotoPostHandler(w http.ResponseWriter, r *http.Request) {
	// Get user session information
	userSession := users.GetUserSession(r)

	// Parse our multipart form, 2 << 20 specifies a maximum upload of 2 MB files
	r.ParseMultipartForm(2 << 20)

	// Get the file and handler from the form
	file, fileHeader, err := r.FormFile("avatar")
	if err != nil {
		log.Println(err)
	}
	defer file.Close()

	// Get Upload Directory
	uploadDir := filestore.GetUploadPath(userSession.Username, "avatar")

	// Upload file to directory
	filepath, err := filestore.Upload(file, fileHeader, uploadDir)
	if err != nil {
		log.Println(err)
		flash.Set(w, r, flash.ERROR, "Profile update failed")
		http.Redirect(w, r, "/profiles/show", http.StatusSeeOther)
		return
	}

	// Save file path to database
	queries := New(database.DB)
	updateProfileParams := UpdateProfileParams{
		Avatar:    sql.NullString{String: filepath, Valid: true},
		UpdatedAt: time.Now().String(),
		UserID:    userSession.Id,
	}
	_, err = queries.UpdateProfile(context.Background(), updateProfileParams)
	if err != nil {
		log.Println(err)
		flash.Set(w, r, flash.ERROR, "Profile update failed")
	} else {
		flash.Set(w, r, flash.SUCCESS, "Profile successfully updated")
	}

	// Redirect to profile page
	http.Redirect(w, r, "/profiles/show", http.StatusSeeOther)
}
