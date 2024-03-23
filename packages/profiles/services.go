package profiles

import (
	"context"
	"database/sql"
	"encoding/base64"
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/larry-lw-chan/goti/database"
	"github.com/larry-lw-chan/goti/packages/filestore"
	"github.com/larry-lw-chan/goti/packages/sessions/flash"
)

/****************************************************************
* Data Handling Logic
****************************************************************/
func createProfileForUser(userID int64) (Profile, error) {
	queries := New(database.DB)

	profileParem := CreateProfileParams{
		Name:      sql.NullString{String: "your name", Valid: true},
		Bio:       sql.NullString{String: "add bio", Valid: true},
		Link:      sql.NullString{String: "add links", Valid: true},
		Private:   sql.NullInt64{Int64: 0, Valid: true},
		Avatar:    sql.NullString{Valid: false},
		UserID:    userID,
		CreatedAt: time.Now().String(),
		UpdatedAt: time.Now().String(),
	}
	profile, err := queries.CreateProfile(context.Background(), profileParem)

	return profile, err
}

func updateProfileForUser(userID int64, name, bio, link string, private int64) (Profile, error) {
	queries := New(database.DB)
	updateProfileParams := UpdateProfileParams{
		Name:      sql.NullString{String: name, Valid: true},
		Bio:       sql.NullString{String: bio, Valid: true},
		Link:      sql.NullString{String: link, Valid: true},
		Private:   sql.NullInt64{Int64: private, Valid: true},
		UpdatedAt: time.Now().String(),
		UserID:    userID,
	}
	// Update profile
	profile, err := queries.UpdateProfile(context.Background(), updateProfileParams)
	return profile, err
}

func saveFilePathToProfile(userID int64, filepath string) (Profile, error) {
	queries := New(database.DB)
	updateProfileParams := UpdateProfileParams{
		Avatar:    sql.NullString{String: filepath, Valid: true},
		UpdatedAt: time.Now().String(),
		UserID:    userID,
	}
	profile, err := queries.UpdateProfile(context.Background(), updateProfileParams)
	return profile, err
}

func decodeBase64Image(base64Image string) ([]byte, string, error) {
	// Decode base64 image
	coI := strings.Index(string(base64Image), ",")
	rawData := string(base64Image)[coI+1:]

	// Decode base64 data and assign name pattern
	fileBytes, err := base64.StdEncoding.DecodeString(string(rawData))
	if err != nil {
		return nil, "", nil
	}
	namePattern := "avatar-*.png" // croppie defaults to png
	return fileBytes, namePattern, err
}

func handleMultipartFile(w http.ResponseWriter, r *http.Request) ([]byte, string) {
	// Get the file and handler from the form
	file, fileHeader, err := r.FormFile("avatar")
	handleError(w, r, err, "/profiles/edit/photo")
	defer file.Close()

	// Print File Header Log
	filestore.PrintFileHeader(fileHeader)

	// Convert File into byte slice
	fileBytes, err := io.ReadAll(file)
	handleError(w, r, err, "/profiles/edit/photo")

	// Get name pattern from file header
	ext := strings.Split(fileHeader.Filename, ".")[1]
	namePattern := "avatar-*." + ext

	return fileBytes, namePattern
}

/****************************************************************
* Helper Functions
****************************************************************/
func handleError(w http.ResponseWriter, r *http.Request, err error, redirect string) {
	if err != nil {
		log.Println(err)
		flash.Set(w, r, flash.ERROR, err.Error())
		http.Redirect(w, r, redirect, http.StatusSeeOther)
		return
	}
}
