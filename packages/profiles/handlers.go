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

// Uploads a photo to the server and redirect to profile avatar edit page
func CreatePhotoHandler(w http.ResponseWriter, r *http.Request) {
	data := r.Context().Value("data").(map[string]interface{})

	// Get user session information
	userSession := users.GetUserSession(r)

	// Parse our multipart form, 2 << 20 specifies a maximum upload of 2 MB files
	r.ParseMultipartForm(2 << 20)

	// Get the file and handler from the form
	file, fileHeader, err := r.FormFile("avatar")
	handleError(w, r, err, "/profiles/edit/photo")

	// Place data in file struct
	fileUpload := filestore.FileUpload{
		File:       file,
		FileHeader: fileHeader,
		Directory:  userSession.Username + "/avatar",
	}

	// Upload file to directory
	filepath, err := filestore.Upload(fileUpload)
	handleError(w, r, err, "/profiles/edit/photo")

	// Save file path to database
	queries := New(database.DB)
	updateProfileParams := UpdateProfileParams{
		Avatar:    sql.NullString{String: filepath, Valid: true},
		UpdatedAt: time.Now().String(),
		UserID:    userSession.Id,
	}
	profile, err := queries.UpdateProfile(context.Background(), updateProfileParams)
	if err != nil {
		flash.Set(w, r, flash.ERROR, "Profile update failed")
	}

	data["Profile"] = profile

	http.Redirect(w, r, "/profiles/edit/photo", http.StatusSeeOther)
}

func EditPhotoHandler(w http.ResponseWriter, r *http.Request) {
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

	render.Template(w, data, "/profiles/edit-photo.app.tmpl")
}

func EditPhotoPostHandler(w http.ResponseWriter, r *http.Request) {
	// Get user session information
	userSession := users.GetUserSession(r)

	// Parse our multipart form, 2 << 20 specifies a maximum upload of 2 MB files
	r.ParseMultipartForm(2 << 20)

	// Get the file and handler from the form
	file, fileHeader, err := r.FormFile("avatar")
	handleError(w, r, err, "/profiles/show")

	// Place data in file struct
	fileUpload := filestore.FileUpload{
		File:       file,
		FileHeader: fileHeader,
		Directory:  userSession.Username + "/avatar",
	}

	// Upload file to directory
	filepath, err := filestore.Upload(fileUpload)
	handleError(w, r, err, "/profiles/show")

	// Save file path to database
	queries := New(database.DB)
	updateProfileParams := UpdateProfileParams{
		Avatar:    sql.NullString{String: filepath, Valid: true},
		UpdatedAt: time.Now().String(),
		UserID:    userSession.Id,
	}
	_, err = queries.UpdateProfile(context.Background(), updateProfileParams)
	if err != nil {
		flash.Set(w, r, flash.ERROR, "Profile update failed")
	} else {
		flash.Set(w, r, flash.SUCCESS, "Profile successfully updated")
	}

	// Redirect to profile page
	http.Redirect(w, r, "/profiles/show", http.StatusSeeOther)
}

func DeletePhotoPostHandler(w http.ResponseWriter, r *http.Request) {
	// Get avatar path
	userSession := users.GetUserSession(r)
	queries := New(database.DB)

	// Get avatar path
	filename, err := queries.GetProfileAvatarFromUserId(context.Background(), userSession.Id)
	handleError(w, r, err, "/profiles/show")

	// Delete file from filestore
	err = filestore.Delete(filename.String)
	handleError(w, r, err, "/profiles/show")

	// Fill in update profile parameters
	updateProfileParams := UpdateProfileParams{
		Avatar: sql.NullString{Valid: false},
		UserID: userSession.Id,
	}

	// Delete file path from database
	_, err = queries.UpdateProfile(context.Background(), updateProfileParams)
	handleError(w, r, err, "/profiles/show")

	// Redirect to profile page
	flash.Set(w, r, flash.SUCCESS, "Profile photo successfully deleted")
	http.Redirect(w, r, "/profiles/show", http.StatusSeeOther)
}

func handleError(w http.ResponseWriter, r *http.Request, err error, redirect string) {
	if err != nil {
		log.Println(err)
		flash.Set(w, r, flash.ERROR, err.Error())
		http.Redirect(w, r, redirect, http.StatusSeeOther)
		return
	}
}
