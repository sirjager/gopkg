package tokens

import "time"

// TokenBuilder is an interface for managing tokens
type TokenBuilder interface {
	// Create Token if token for specific duration
	CreateToken(data interface{}, tokenExpiresIn time.Duration) (string, *Payload, error)

	// Validates token integrity and expiration time
	VerifyToken(token string, data interface{}) (*Payload, error)

	// ReadPayload extracts custom payload data from payload if valid
	ReadPayload(payload *Payload, data interface{}) error
}
