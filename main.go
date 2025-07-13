package main

import (
	"net/http"

	"github.com/SanketJawali/url-shortner-api/routes"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	// Init router
	r := chi.NewRouter()

	// Define routes
	r.Use(middleware.Logger)
	r.Get("/hello", routes.HelloHandler)
	r.Get("/auth/register", routes.HelloHandler)
	r.Get("/auth/login", routes.HelloHandler)
	r.Get("/auth/verify", routes.HelloHandler)

	// Run the server
	http.ListenAndServe(":8080", r)
}
