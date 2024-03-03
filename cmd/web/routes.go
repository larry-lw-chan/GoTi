package main

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/larry-lw-chan/go-tiny/packages/pages"
	"github.com/larry-lw-chan/go-tiny/packages/users"
)

func Routes() *chi.Mux {
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

	// Routes
	r.Get("/", pages.HomePageHandler)
	r.Get("/about", pages.AboutPageHandler)
	r.Get("/contact", pages.ContactPageHandler)
	r.Get("/privacy", pages.PrivacyPageHandler)

	// Authentication Routes
	r.Get("/login", users.LoginHandler)
	r.Get("/register", users.RegisterHandler)
	r.Post("/register", users.RegisterPostHandler)
	r.Get("/forgot-password", users.ForgotPasswordHandler)

	// Return the Router
	return r
}
