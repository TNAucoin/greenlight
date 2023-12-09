package main

import (
	"fmt"
	"github.com/tnaucoin/greenlight/internal/data"
	"github.com/tnaucoin/greenlight/internal/validator"
	"net/http"
	"time"
)

func (app *application) createMovieHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Title   string       `json:"title"`
		Year    int32        `json:"year"`
		Runtime data.Runtime `json:"runtime"`
		Genres  []string     `json:"genres"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	v := validator.New()
	// Validation checks for creating a movie record.
	v.Check(input.Title != "", "title", "this field cannot be blank")
	v.Check(len(input.Title) < 500, "title", "this field is too long (maximum is 500 characters)")
	v.Check(input.Year != 0, "year", "this field cannot be blank")
	v.Check(input.Year >= 1888, "year", "must be greater than 1888")
	v.Check(input.Year <= int32(time.Now().Year()), "year", "must not be in the future")
	v.Check(input.Runtime != 0, "runtime", "this field cannot be blank")
	v.Check(input.Runtime > 0, "runtime", "must be a positive integer")
	v.Check(input.Genres != nil, "genres", "this field cannot be blank")
	v.Check(len(input.Genres) >= 1, "genres", "you must specify at least one genre")
	v.Check(len(input.Genres) <= 5, "genres", "you cannot specify more than five genres")
	v.Check(validator.Unique(input.Genres), "genres", "you cannot specify duplicate genres")

	if !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	fmt.Fprintf(w, "%+v\n", input)
}

func (app *application) showMoviesHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil || id < 1 {
		app.notFoundResponse(w, r)
		return
	}
	m := data.Movie{
		ID:        id,
		CreatedAt: time.Now(),
		Title:     "Casablanca",
		Runtime:   102,
		Genres:    []string{"drama", "romance"},
		Version:   1,
	}
	err = app.writeJSON(w, http.StatusOK, envelope{"movie": m}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}

}
