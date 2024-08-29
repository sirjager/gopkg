package httpx

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/rs/zerolog"
)

func Logger(logr zerolog.Logger, disableColors ...bool) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now() // Start timer
			path := r.URL.Path

			rec := &ResponseRecorder{ResponseWriter: w, StatusCode: 200, Body: &bytes.Buffer{}}
			next.ServeHTTP(rec, r)

			duration := time.Since(start)

			event := logr.Info()

			if rec.StatusCode != http.StatusOK {
				var data map[string]interface{}
				if err := json.Unmarshal(rec.Body.Bytes(), &data); err != nil {
					data = map[string]interface{}{}
				}
				event = logr.Error().Interface("error", data["message"])
			}

			if rec.StatusCode >= 400 && rec.StatusCode < 500 {
				event = logr.Warn()
			} else if rec.StatusCode >= 500 {
				event = logr.Error()
			}

			shortenedPath := shortenPath(path, 20)
			icon := getIcon(rec.StatusCode)
			coloredIcon := getColoredIcon(rec.StatusCode)

			event.
				Str("method", r.Method).
				Str("path", shortenedPath).
				Dur("latency", duration).
				Int("code", rec.StatusCode)

			if len(disableColors) > 0 && disableColors[0] {
				event.Msg(icon)
			} else {
				event.Msg(coloredIcon)
			}
		})
	}
}

const (
	boldGreen  = "\033[1;32m"
	boldRed    = "\033[1;31m"
	boldYellow = "\033[1;33m"
	boldCyan   = "\033[1;36m"
	reset      = "\033[0m"
)

func getIcon(code int) string {
	switch code {
	case 200:
		return "" // OK
	case 201:
		return "" // Created
	case 204:
		return "" // No Content
	case 301:
		return "" // Moved Permanently
	case 302:
		return "" // Found
	case 304:
		return "" // Not Modified
	case 400:
		return "" // Bad Request
	case 401:
		return "" // Unauthorized
	case 403:
		return "" // Forbidden
	case 404:
		return "" // Not Found
	case 500:
		return "" // Internal Server Error
	case 502:
		return "" // Bad Gateway
	case 503:
		return "" // Service Unavailable
	default:
		return "" // Unknown or unhandled status
	}
}

func getColoredIcon(code int) string {
	switch code {
	case 200:
		return fmt.Sprintf("%s %s", boldGreen, reset)
	case 201:
		return fmt.Sprintf("%s %s", boldGreen, reset)
	case 204:
		return fmt.Sprintf("%s %s", boldGreen, reset)
	case 301:
		return fmt.Sprintf("%s %s", boldCyan, reset)
	case 302:
		return fmt.Sprintf("%s %s", boldCyan, reset)
	case 304:
		return fmt.Sprintf("%s %s", boldCyan, reset)
	case 400:
		return fmt.Sprintf("%s %s", boldYellow, reset)
	case 401:
		return fmt.Sprintf("%s %s", boldYellow, reset)
	case 403:
		return fmt.Sprintf("%s %s", boldYellow, reset)
	case 404:
		return fmt.Sprintf("%s %s", boldYellow, reset)
	case 500:
		return fmt.Sprintf("%s %s", boldRed, reset)
	case 502:
		return fmt.Sprintf("%s %s", boldRed, reset)
	case 503:
		return fmt.Sprintf("%s %s", boldRed, reset)
	default:
		return fmt.Sprintf("%s %s", boldRed, reset)
	}
}

func shortenPath(path string, max int) string {
	if len(path) > max {
		return path[0:max] + "..."
	}
	return path
}

type ResponseRecorder struct {
	http.ResponseWriter
	Body       *bytes.Buffer
	StatusCode int
}

func (rec *ResponseRecorder) WriteHeader(statusCode int) {
	rec.StatusCode = statusCode
	rec.ResponseWriter.WriteHeader(statusCode)
}

func (rec *ResponseRecorder) Write(b []byte) (int, error) {
	rec.Body.Write(b)
	return rec.ResponseWriter.Write(b)
}
