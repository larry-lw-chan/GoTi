package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/larry-lw-chan/go-tiny/handlers"
)

func main() {
	// Define Routes Here
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", handlers.HomePage)
	r.Get("/about", handlers.AboutPage)

	// Start Server
	fmt.Println("Server is running at :3000")
	http.ListenAndServe(":3000", r)
}
