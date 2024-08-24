package httpx

import (
	"encoding/json"
	"net/http"
)

// FailureText sends a plain text error response to the client.
// It accepts a `http.ResponseWriter`, an `error` object, and an optional `statusCode`.
// The default status code is 500 unless overridden by the optional `statusCode`.
// The error message is written as the response body in plain text format.
// If writing the response fails, a plain text error message is sent with a status code of 500.
func FailureText(w http.ResponseWriter, err error, statusCode ...int) {
	status := 500
	if len(statusCode) > 0 {
		status = statusCode[0]
	}
	w.Header().Add("Content-Type", "text/plain")
	w.WriteHeader(status)
	if _, writeErr := w.Write([]byte(err.Error())); writeErr != nil {
		http.Error(w, "failed to write response message", http.StatusInternalServerError)
	}
}

// SuccessText sends a plain text success response to the client.
// It accepts a `http.ResponseWriter`, a `message` string, and an optional `statusCode`.
// The default status code is 200 unless overridden by the optional `statusCode`.
// The `message` is written as the response body in plain text format.
// If writing the response fails, a plain text error message is sent with a status code of 500.
func SuccessText(w http.ResponseWriter, message string, statusCode ...int) {
	status := 200
	if len(statusCode) > 0 {
		status = statusCode[0]
	}
	w.Header().Add("Content-Type", "text/plain")
	w.WriteHeader(status)
	if _, err := w.Write([]byte(message)); err != nil {
		http.Error(w, "failed to write response message", http.StatusInternalServerError)
	}
}

// SuccessJSON sends a successful JSON response to the client.
// It accepts a `http.ResponseWriter`, a `response` object, and an optional `statusCode`.
// The default status code is 200 unless overridden by the optional `statusCode`.
// The `response` is serialized to JSON and written to the response body.
func Success(w http.ResponseWriter, response any, statusCode ...int) {
	responseJSON(w, response, statusCode...)
}

// ErrorJSON sends an error JSON response to the client.
// It accepts a `http.ResponseWriter`, an `error` object, and an optional `statusCode`.
// The default status code for errors is 500 unless overridden by the optional `statusCode`.
// The `error` message is wrapped in a JSON object and sent as the response body.
func Error(w http.ResponseWriter, err error, statusCode ...int) {
	responseJSON(w, err, statusCode...)
}

// ResponseJSON sends a JSON response to the client.
// It accepts a `http.ResponseWriter`, a `response` object, and an optional `statusCode`.
// If the `response` is an error, it will wrap the error message in a JSON object.
// The `statusCode` can be provided to override the default status code (200).
// The `response` is serialized to JSON and written to the response body.
func responseJSON(w http.ResponseWriter, response any, statusCode ...int) {
	status := 200
	if err, isErr := response.(error); isErr {
		status = 500
		response = map[string]string{"error": err.Error()}
	}
	if len(statusCode) > 0 {
		status = statusCode[0]
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		// This ensures that your error response is also formatted as JSON.
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, `{"error": "Failed to encode response"}`, http.StatusInternalServerError)
	}
}
