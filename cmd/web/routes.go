package main

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/larry-lw-chan/goti/internal/activities"
	"github.com/larry-lw-chan/goti/internal/auth"
	"github.com/larry-lw-chan/goti/internal/pages"
	"github.com/larry-lw-chan/goti/internal/profiles"
	"github.com/larry-lw-chan/goti/internal/search"
	"github.com/larry-lw-chan/goti/internal/threads"
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

	// Asset File Server
	assetFS := http.FileServer(http.Dir(tmplPath + "/assets"))
	r.Handle("/assets/*", http.StripPrefix("/assets/", assetFS))

	// Local Upload File Server
	uploadFS := http.FileServer(http.Dir("uploads"))
	r.Handle("/uploads/*", http.StripPrefix("/uploads/", uploadFS))

	// Register Package Routes
	r.Mount("/", pages.Router())
	r.Mount("/auth", auth.Router())
	r.Mount("/profiles", profiles.Router())
	r.Mount("/threads", threads.Router())
	r.Mount("/search", search.Router())
	r.Mount("/activities", activities.Router())

	// Custom Error Handling
	r.NotFound(pages.NotFoundHandler)

	// Return the Router
	return r
}
