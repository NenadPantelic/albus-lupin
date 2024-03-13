package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func (app *application) routes() http.Handler {
	mux := chi.NewRouter()

	// if the app panics, this will catch it, recover, log the stacktrace and return 500 status code, the app will go on
	mux.Use(middleware.Recoverer)

	mux.Get("/status", app.Status)
	mux.Post("/grade", app.Grade)

	// mux.Route("/admin", func(mux chi.Router) {
	// 	mux.Use(app.authRequired)

	// 	mux.Get("/movies", app.MovieCatalog)
	// })

	return mux
}
