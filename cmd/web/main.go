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
	"github.com/larry-lw-chan/goti/packages/auth"
	"github.com/larry-lw-chan/goti/packages/cookie"
	"github.com/larry-lw-chan/goti/packages/pages"
	"github.com/larry-lw-chan/goti/packages/users"
	"github.com/larry-lw-chan/goti/packages/utils/render"
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
	r.Mount("/", pages.Router())
	r.Mount("/auth", auth.Router())
	r.Mount("/users", users.Router())

	// Return the Router
	return r
}

func loadTemplates() {
	// Define the path to the templates
	path := "./templates/default"

	// Override the default template path with user configuration
	// if os.Getenv("TEMPLATE_PATH") != "" {
	// 	path = os.Getenv("TEMPLATE_PATH")
	// 	log.Println(path)
	// }

	// Setup Template Layouts
	render.Layouts(render.Location{
		TmplPath: path,
		Layout: map[string][]string{
			"base": {
				"/layout/base.tmpl",
			},
			"app": {
				"/layout/app.tmpl",
			},
		},
	})
}

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	return ":" + port
}

func init() {
	// Load .env configuration
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	// Load Templates
	loadTemplates()

	// Initialize session store
	cookie.InitializeStore()

	// Connect to the database
	db := data.Connect()
	defer db.Close() // Defer close the database connection

	// Load Routes
	r := routes()

	// Start Server
	port := getPort()
	fmt.Println("Server is running at " + port)
	http.ListenAndServe(port, r)
}
