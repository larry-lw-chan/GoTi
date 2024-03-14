package users

import (
	"context"
	"log"

	"github.com/larry-lw-chan/goti/database"
	"golang.org/x/crypto/bcrypt"
)

/*************************************************
* Authenticate User
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

/*************************************************
* HASHING SERVICES
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
