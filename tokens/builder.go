package tokens

import "time"

// TokenBuilder is an interface for managing tokens
type TokenBuilder interface {
	// Create Token if token for specific duration
	CreateToken(payloadData interface{}, tokenAliveDuration time.Duration) (string, *Payload, error)

	// Validates token integrity and expiration time
	VerifyToken(token string) (*Payload, error)
}
