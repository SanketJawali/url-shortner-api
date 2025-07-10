package main

import (
	"net/http"

	"github.com/SanketJawali/url-shortner-api/routes"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Get("/", routes.HelloHandler)
	http.ListenAndServe(":8080", r)
}
