package httpx

import (
	"net/http"
)

// Redirect creates middleware that redirects all incoming requests
// to the provided URL with the specified HTTP status code.
//
// Parameters:
// - to: The destination URL to redirect to.
// - code: The HTTP status code to use for the redirection (e.g., 301, 302).
//
// Returns:
// - A middleware function that handles the redirection.
func Redirect(to string, code int) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Redirect(w, r, to, code)
		})
	}
}

// RedirectRoute creates a handler function that redirects incoming requests
// to the provided URL with the specified HTTP status code.
//
// Parameters:
// - to: The destination URL to redirect to.
// - code: The HTTP status code to use for the redirection (e.g., 301, 302).
//
// Returns:
// - A handler function that performs the redirection.
func RedirectRoute(to string, code int) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, to, code)
	}
}
