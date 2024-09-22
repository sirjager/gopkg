package httpx

import (
	"net/http"
)

// Redirect redirects to the provided url
func Redirect(to string, code int) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Redirect(w, r, to, code)
		})
	}
}
