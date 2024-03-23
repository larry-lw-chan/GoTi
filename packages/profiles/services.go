package profiles

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"time"

	"github.com/larry-lw-chan/goti/database"
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
