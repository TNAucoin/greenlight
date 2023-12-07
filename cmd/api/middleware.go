package main

import (
	"fmt"
	"net/http"
)

// recoverPanic is a middleware that recovers from panics, logs the error using our
// custom Logger type, and returns a '500 Internal Server Error' response if
// possible.
func (app *application) recoverPanic(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			// Use the builtin recover function to check if there has been a panic or
			// not.
			if err := recover(); err != nil {
				// If there was a panic set Connection:close to tell go http server to
				// close the current connection after a response has been sent.
				w.Header().Set("Connection", "close")
				// Call the app.serverErrorResponse() method to send a 500 Internal
				// Server Error response containing a generic error message.
				// The value returned from the recover() function is of the empty
				// interface type, so we use fmt.Errorf() to normalize it into an
				// error.
				app.serverErrorResponse(w, r, fmt.Errorf("%s", err))
			}
		}()
		next.ServeHTTP(w, r)
	})
}
