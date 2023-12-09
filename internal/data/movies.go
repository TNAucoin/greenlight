package data

import (
	"github.com/tnaucoin/greenlight/internal/validator"
	"time"
)

type Movie struct {
	ID        int64
	CreatedAt time.Time
	Title     string
	Year      int32
	Runtime   Runtime
	Genres    []string
	Version   int32
}

// validateTitle checks the title field for the movie is not blank and is not
func validateTitle(v *validator.Validator, title string) {
	v.Check(title != "", "title", "this field cannot be blank")
	v.Check(len(title) < 500, "title", "this field is too long (maximum is 500 characters)")
}

func validateYear(v *validator.Validator, year int32) {
	v.Check(year != 0, "year", "this field cannot be blank")
	v.Check(year >= 1888, "year", "must be greater than 1888")
	v.Check(year <= int32(time.Now().Year()), "year", "must not be in the future")
}

func validateRuntime(v *validator.Validator, runtime Runtime) {
	v.Check(runtime != 0, "runtime", "this field cannot be blank")
	v.Check(runtime > 0, "runtime", "must be a positive integer")
}

func validateGenres(v *validator.Validator, genres []string) {
	v.Check(genres != nil, "genres", "this field cannot be blank")
	v.Check(len(genres) >= 1, "genres", "you must specify at least one genre")
	v.Check(len(genres) <= 5, "genres", "you cannot specify more than five genres")
	v.Check(validator.Unique(genres), "genres", "you cannot specify duplicate genres")
}

func ValidateMovie(v *validator.Validator, m *Movie) {
	validateTitle(v, m.Title)
	validateYear(v, m.Year)
	validateRuntime(v, m.Runtime)
	validateGenres(v, m.Genres)
}
