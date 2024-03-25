package main

import (
	"encoding/base64"
	"fmt"

	"github.com/gorilla/securecookie"
)

func main() {
	authKey := securecookie.GenerateRandomKey(64)
	encryptionKey := securecookie.GenerateRandomKey(32)

	authKeyStr := base64.StdEncoding.EncodeToString(authKey)
	encryptionKeyStr := base64.StdEncoding.EncodeToString(encryptionKey)

	fmt.Print("\n")
	fmt.Print("\n")
	fmt.Print("\n")
	fmt.Print("AUTH_KEY=" + authKeyStr + "\n")
	fmt.Print("ENCRYPTION_KEY=" + encryptionKeyStr + "\n")
}
