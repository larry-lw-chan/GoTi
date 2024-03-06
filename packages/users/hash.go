package users

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

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
