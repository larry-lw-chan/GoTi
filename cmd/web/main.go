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
	"github.com/larry-lw-chan/goti/database"
	"github.com/larry-lw-chan/goti/packages/pages"
	"github.com/larry-lw-chan/goti/packages/sessions/cookie"
	"github.com/larry-lw-chan/goti/packages/sessions/flash"
	"github.com/larry-lw-chan/goti/packages/users"
	"github.com/larry-lw-chan/goti/packages/utils/render"
)

// Define the path to the templates
var path string = "templates/default"
var port string = ":3000"

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

	// Load Flash Middleware
	r.Use(flash.TryGetFlash)

	// Asset File Server
	assetFS := http.FileServer(http.Dir(path + "/assets"))
	r.Handle("/assets/*", http.StripPrefix("/assets/", assetFS))

	// Local Upload File Server
	uploadFS := http.FileServer(http.Dir("uploads"))
	r.Handle("/uploads/*", http.StripPrefix("/uploads/", uploadFS))

	// Register Package Routes
	r.Mount("/", pages.Router())
	r.Mount("/auth", users.AuthRouter())
	r.Mount("/users", users.UserRouter())

	// Return the Router
	return r
}

// Inject Template Layouts
func loadTemplates() {
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

// Load Envrionment Configuration
func init() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Override the default template path with user configuration
	if os.Getenv("TEMPLATE_PATH") != "" {
		path = os.Getenv("TEMPLATE_PATH")
	}

	// Override the default port with user configuration
	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}
}

func main() {
	// Load Templates
	loadTemplates()

	// Initialize session store
	cookie.InitializeStore()

	// Connect to the database
	db := database.Connect()
	defer db.Close() // Defer close the database connection

	// Load Routes
	r := routes()

	// Start Server
	fmt.Println("Server is running at " + port)
	http.ListenAndServe(port, r)
}
