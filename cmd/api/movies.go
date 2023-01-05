package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

const movie_id = "id"

func (app *application) showMovieHandler(w http.ResponseWriter, r *http.Request) {
	idString := chi.URLParam(r, movie_id)

	id, err := strconv.ParseInt(idString, 10, 64)
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "show movie #%d", id)
}

func (app *application) createMovieHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "create a new movie")
}
