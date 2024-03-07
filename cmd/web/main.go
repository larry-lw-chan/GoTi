package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
	"github.com/larry-lw-chan/goti/data"
	"github.com/larry-lw-chan/goti/packages/pages"
	"github.com/larry-lw-chan/goti/packages/users"
	_ "github.com/mattn/go-sqlite3"
)

func routes() *chi.Mux {
	// Define Routes Here
	r := chi.NewRouter()

	// A good base middleware stack (Chi Default)
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	r.Use(middleware.Timeout(60 * time.Second))

	// Static File Server
	fs := http.FileServer(http.Dir("static"))
	r.Handle("/static/*", http.StripPrefix("/static/", fs))

	// Register Package Routes
	pages.Routes(r)
	users.Routes(r)

	// Return the Router
	return r
}

func main() {
	// Load config from environment
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db := data.Connect() // Connect Database
	defer db.Close()     // Defer close the database connection

	// Load Routes
	r := routes()

	// Start Server
	if os.Getenv("PORT") == "" {
		os.Setenv("PORT", "3000")
	}
	port := ":" + os.Getenv("PORT")

	fmt.Println("Server is running at " + port)
	http.ListenAndServe(port, r)
}
