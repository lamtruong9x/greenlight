package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (app *application) routes() http.Handler {
	router := chi.NewRouter()

	// set custom method not allowed and not found handler
	router.MethodNotAllowed(app.methodNotAllowedResponse)
	router.NotFound(app.notFoundResponse)

	// set routes
	router.Route("/v1", func(r chi.Router) {
		r.Get("/healthcheck", app.healthcheckHandler)
		r.Post("/movies", app.createMovieHandler)
		r.Get("/movies/{id}", app.showMovieHandler)
	})
	return router
}
