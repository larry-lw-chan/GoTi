package auth

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
)

func RandBase64String(length int) string {
	b := make([]byte, length)

	_, err := rand.Read(b)
	if err != nil {
		fmt.Println("error:", err)
		return err.Error()
	}

	return base64.StdEncoding.EncodeToString(b)
}
