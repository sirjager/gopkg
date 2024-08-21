package utils

import (
	"encoding/base64"
	"strings"
)

// BytesToBase64 converts bytes to base64 string.
// This ensures that the resulting string can be safely used in text-based contexts like cookies, URLs, and HTTP headers,
// which may not handle arbitrary binary data properly.
func BytesToBase64(value []byte) string {
	return base64.StdEncoding.EncodeToString(value)
}

// Base64ToBytes converts string to base64
func Base64ToBytes(value string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(value)
}

type StringBuilder struct {
	builder strings.Builder
}

// Write appends a string to the StringBuilder
func (sb *StringBuilder) Write(s string) {
	sb.builder.WriteString(s)
}

// WriteLine appends a string followed by a newline to the StringBuilder
func (sb *StringBuilder) WriteLine(s string) {
	sb.builder.WriteString(s)
	sb.builder.WriteString("\n")
}

// String returns the complete string built by the StringBuilder
func (sb *StringBuilder) String() string {
	return sb.builder.String()
}

// Reset clears the StringBuilder for reuse
func (sb *StringBuilder) Reset() {
	sb.builder.Reset()
}
