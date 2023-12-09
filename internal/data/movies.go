package data

import (
	"github.com/tnaucoin/greenlight/internal/validator"
	"time"
)

type Movie struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"-"`
	Title     string    `json:"title"`
	Year      int32     `json:"year,omitempty"`
	Runtime   Runtime   `json:"runtime,omitempty"`
	Genres    []string  `json:"genres,omitempty"`
	Version   int32     `json:"version"`
}

func ValidateMovie(v *validator.Validator, m *Movie) {
	v.Check(m.Title != "", "title", "this field cannot be blank")
	v.Check(len(m.Title) < 500, "title", "this field is too long (maximum is 500 characters)")
	v.Check(m.Year != 0, "year", "this field cannot be blank")
	v.Check(m.Year >= 1888, "year", "must be greater than 1888")
	v.Check(m.Year <= int32(time.Now().Year()), "year", "must not be in the future")
	v.Check(m.Runtime != 0, "runtime", "this field cannot be blank")
	v.Check(m.Runtime > 0, "runtime", "must be a positive integer")
	v.Check(m.Genres != nil, "genres", "this field cannot be blank")
	v.Check(len(m.Genres) >= 1, "genres", "you must specify at least one genre")
	v.Check(len(m.Genres) <= 5, "genres", "you cannot specify more than five genres")
	v.Check(validator.Unique(m.Genres), "genres", "you cannot specify duplicate genres")
}
