package main

import (
	"fmt"
	"net/http"

	"github.com/larry-lw-chan/goti/database"
	"github.com/larry-lw-chan/goti/internal/sessions/cookie"
	"github.com/larry-lw-chan/goti/internal/utils/render"
)

// Global Defaults
type Config struct {
	TmplPath      string
	LayoutFolder  string
	Port          string
	AuthKey       string
	EncryptionKey string
	MaxAge        string
}

func main() {
	// Set Default Config
	c := Config{
		TmplPath:     "templates/default",
		LayoutFolder: "layout",
		Port:         ":8080",
	}

	// Load User ENV Configurations
	loadEnvConfig(&c)

	// Load Templates
	render.New(render.Options{
		TmplPath:     c.TmplPath,
		LayoutFolder: c.LayoutFolder,
	})

	// Initialize session store
	cookie.NewStore(cookie.Options{
		AuthKey:       c.AuthKey,
		EncryptionKey: c.EncryptionKey,
		MaxAge:        c.MaxAge,
	})

	// Connect to the database
	db := database.Connect()
	defer db.Close() // Defer close the database connection

	// Load Routes
	r := routes(&c)

	// Start Server
	fmt.Println("Server is running at " + c.Port)
	http.ListenAndServe(c.Port, r)
}
