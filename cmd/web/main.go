package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/larry-lw-chan/go-tiny/packages/handlers"
)

func main() {
	// Define Routes Here
	r := chi.NewRouter()

	// Middleware
	r.Use(middleware.Logger)

	// Static File Server
	fs := http.FileServer(http.Dir("static"))
	r.Handle("/static/*", http.StripPrefix("/static/", fs))

	// Routes
	r.Get("/", handlers.HomePage)
	r.Get("/about", handlers.AboutPage)
	r.Get("/contact", handlers.ContactPage)

	// Start Server
	fmt.Println("Server is running at :3000")
	http.ListenAndServe(":3000", r)
}
