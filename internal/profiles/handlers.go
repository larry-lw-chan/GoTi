package profiles

import (
	"context"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/larry-lw-chan/goti/database"
	"github.com/larry-lw-chan/goti/internal/auth"
	"github.com/larry-lw-chan/goti/internal/filestore"
	"github.com/larry-lw-chan/goti/internal/helper"
	"github.com/larry-lw-chan/goti/internal/sessions/flash"
	"github.com/larry-lw-chan/goti/internal/utils/render"
)

/****************************************************************
* Profile Handlers
****************************************************************/

func ShowHandler(w http.ResponseWriter, r *http.Request) {
	data := r.Context().Value("data").(map[string]interface{})

	// Get username from chi and session information
	username := chi.URLParam(r, "username")
	userSession := auth.GetUserSession(r)
	queries := New(database.DB)

	// Get profile information by username or user session if username is empty
	var profile Profile
	var err error
	if username != "" {
		profile, err = queries.GetProfileFromUsername(context.Background(), username)
		if err != nil {
			helper.PageNotFound(w, http.StatusNotFound)
			return
		}
	} else {
		profile, err = queries.GetProfileFromUserId(context.Background(), userSession.ID)
		if err != nil {
			flash.Set(w, r, flash.NOTICE, "Please complete your user profile")
			http.Redirect(w, r, "/profiles/edit", http.StatusSeeOther)
			return
		}
	}

	// Load username and profile to pass to templateUsername
	data["Profile"] = profile

	// Check if profile belongs to user
	if profile.UserID == userSession.ID {
		data["EditProfile"] = true
	} else {
		data["EditProfile"] = false
	}

	// Load user into profile
	render.Template(w, data,
		"/profiles/show.app.tmpl",
		"/partials/__avatar.app.tmpl")
}

func EditHandler(w http.ResponseWriter, r *http.Request) {
	data := r.Context().Value("data").(map[string]interface{})

	// Get user session information
	userSession := auth.GetUserSession(r)
	queries := New(database.DB)

	// Get profile information
	profile, err := queries.GetProfileFromUserId(context.Background(), userSession.ID)
	if err != nil {
		log.Println(err)
	}

	// Load username and profile to pass to template
	data["Profile"] = profile

	// Load user into profile
	render.Template(w, data, "/profiles/edit.app.tmpl")
}

func EditPostHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	// Get user session information
	userSession := auth.GetUserSession(r)

	// Assign private to (1) if checkbox is checked ("private")
	var private int64 = 0
	if r.FormValue("private") == "private" {
		private = 1
	}

	// Serverside Form Validation
	createProfileValidation := CreateProfileValidation{
		Username: r.FormValue("username"),
		Name:     r.FormValue("name"),
		Bio:      r.FormValue("bio"),
		Link:     r.FormValue("link"),
		Private:  private,
	}

	errs := validateCreateProfile(createProfileValidation)

	if errs != nil {
		messages := helper.GetErrorMessages(errs)
		flash.Set(w, r, flash.ERROR, messages)
		http.Redirect(w, r, "/profiles/edit", http.StatusSeeOther)
		return
	}

	// Create or Update Profile
	username, name, bio, link := r.FormValue("username"), r.FormValue("name"), r.FormValue("bio"), r.FormValue("link")
	_, err := updateOrCreateUserProfile(username, name, bio, link, private, userSession.ID)

	if err != nil {
		log.Println(err)
		flash.Set(w, r, flash.ERROR, "Profile update failed")
	} else {
		flash.Set(w, r, flash.SUCCESS, "Profile successfully updated")
	}

	// Redirect to profile page
	http.Redirect(w, r, "/profiles/show", http.StatusSeeOther)
}

func ShowProfileHandler(w http.ResponseWriter, r *http.Request) {
	data := r.Context().Value("data").(map[string]interface{})
	username := chi.URLParam(r, "username") // Get user session information from chi

	// Get profile information
	queries := New(database.DB)
	profile, err := queries.GetProfileFromUsername(context.Background(), username)
	if err != nil {
		log.Println(err)
	}
	data["Profile"] = profile

	// Load username and profile to pass to template
	render.Template(w, data, "/profiles/show-profile.app.tmpl")
}

/****************************************************************
* Profile Photo Handlers
****************************************************************/
func EditPhotoHandler(w http.ResponseWriter, r *http.Request) {
	data := r.Context().Value("data").(map[string]interface{})

	// Get user session information
	userSession := auth.GetUserSession(r)
	queries := New(database.DB)

	// Get profile information
	profile, err := queries.GetProfileFromUserId(context.Background(), userSession.ID)
	if err != nil {
		log.Println(err)
	}

	// Load profile to pass to template
	data["Profile"] = profile

	render.Template(w, data, "/profiles/edit-photo.app.tmpl")
}

func EditPhotoPostHandler(w http.ResponseWriter, r *http.Request) {
	// Get user session information
	userSession := auth.GetUserSession(r)

	// Parse our multipart form, 2 << 20 specifies a maximum upload of 2 MB files
	r.ParseMultipartForm(2 << 20)

	// Determine whether file is uploaded via base64 or regular file upload
	var fileBytes []byte
	var namePattern string
	var err error

	// Choose between base64 or file upload depending on form input avatar_base64
	if r.FormValue("avatar_base64") != "" {
		Base64Image := r.FormValue("avatar_base64")
		fileBytes, namePattern, err = decodeBase64Image(Base64Image)
		handleError(w, r, err, "/profiles/edit/photo")
	} else {
		fileBytes, namePattern = handleMultipartFile(w, r)
	}

	// Populate File Struct
	fileUpload := filestore.FileUpload{
		FileBytes:   fileBytes,
		NamePattern: namePattern,
		Directory:   userSession.Uuid + "/avatar",
	}

	// Upload file to directory
	filepath, err := filestore.Upload(&fileUpload)
	handleError(w, r, err, "/profiles/edit/photo")

	// Save file path to database
	_, err = createAvatarFilePath(userSession.ID, filepath)

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
	userSession := auth.GetUserSession(r)
	queries := New(database.DB)

	// Get avatar path
	filename, err := queries.GetProfileAvatarFromUserId(context.Background(), userSession.ID)
	handleError(w, r, err, "/profiles/show")

	// Delete file from filestore
	err = filestore.Delete(filename.String)
	handleError(w, r, err, "/profiles/show")

	// Fill in update profile parameters
	_, err = deleteAvatarFilePath(userSession.ID)
	handleError(w, r, err, "/profiles/show")

	// Redirect to profile page
	flash.Set(w, r, flash.SUCCESS, "Profile photo successfully deleted")
	http.Redirect(w, r, "/profiles/show", http.StatusSeeOther)
}
