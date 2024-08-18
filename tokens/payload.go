package tokens

import (
	"errors"
	"time"

	"github.com/sirjager/gopkg/codec"
	"github.com/sirjager/gopkg/utils"
)

var (
	// ErrExpiredToken is returned when a token has expired
	ErrExpiredToken = errors.New("expired token")

	// ErrInvalidToken is returned when a token is invalid
	ErrInvalidToken = errors.New("invalid token")
)

// Payload contains the payload data of the token
type Payload struct {
	IssuedAt  time.Time `json:"iat,omitempty"`
	ExpiresAt time.Time `json:"expires,omitempty"`
	ID        string    `json:"id,omitempty"`
	Data      []byte    `json:"payload,omitempty"`
}

// newPayload creates a new payload for a specific username and duration
func newPayload(
	data interface{},
	duration time.Duration,
	codec codec.Codec,
) (*Payload, error) {
	bytes, err := codec.Marshal(data)
	if err != nil {
		return nil, err
	}
	payload := &Payload{
		Data:      bytes,
		IssuedAt:  time.Now(),
		ID:        utils.XIDNew().String(),
		ExpiresAt: time.Now().Add(duration),
	}
	return payload, nil
}

// Valid checks if the token payload is not expired
func (payload *Payload) Valid() error {
	if time.Now().After(payload.ExpiresAt) {
		return ErrExpiredToken
	}
	return nil
}
