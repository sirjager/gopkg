package httpx

import "net/http"

// SetCookies sets multiple cookies in the HTTP response.
// It accepts a variable number of `*http.Cookie` arguments and sets each cookie
// on the provided `http.ResponseWriter`.
func SetCookies(w http.ResponseWriter, cookies ...*http.Cookie) {
	for _, cookie := range cookies {
		http.SetCookie(w, cookie)
	}
}
