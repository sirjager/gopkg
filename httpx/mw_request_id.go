package httpx

import (
	"context"
	"net/http"

	"github.com/google/uuid"
)

type RequestIDKey int

const (
	ContextRequestID RequestIDKey = iota
)

// RequestID attaches unique request id to each request
func RequestID() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			requestID, err := uuid.NewRandom()
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			ctx := context.WithValue(r.Context(), ContextRequestID, requestID.String())
			w.Header().Set("X-Request-ID", requestID.String())
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
