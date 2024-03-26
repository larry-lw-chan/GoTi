package auth

import (
	"context"
	"log"
	"time"

	"github.com/larry-lw-chan/goti/database"
	"golang.org/x/crypto/bcrypt"
)

/*************************************************
* Authentication
*************************************************/
func AuthenticateUser(email, password string) (*User, bool) {
	// Find user by email
	Queries := New(database.DB)
	user, err := Queries.GetUserFromEmail(context.Background(), email)
	if err != nil {
		return nil, false
	}

	// Check if provided password matches
	if !CheckPasswordHash(password, user.Password) {
		return nil, false
	}
	return &user, true
}

func CreateNewUser(email, password string) (*User, error) {
	hashPwd := HashPassword([]byte(password)) // Hash Password

	queries := New(database.DB)
	createUser := CreateUserParams{
		Uuid:      generateUUID(),
		Email:     email,
		Password:  hashPwd,
		CreatedAt: time.Now().String(),
		UpdatedAt: time.Now().String(),
	}
	user, err := queries.CreateUser(context.Background(), createUser)

	return &user, err
}

/*************************************************
* Hashing Services
*************************************************/
func HashPassword(pwd []byte) string {
	bytes, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	return string(bytes)
}

func CheckPasswordHash(pwd, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pwd))
	return err == nil
}

/************************************************************
* Helper Methods
***********************************************************/
func getErrorMessages(errs []error) (message string) {
	for _, err := range errs {
		message += err.Error() + "<br />"
	}
	return message
}
