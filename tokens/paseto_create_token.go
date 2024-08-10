package tokens

import (
	"time"
)

// CreateToken creates a new token
func (builder *PasetoBuilder) CreateToken(
	payloadData interface{},
	tokenAliveDuration time.Duration,
) (string, *Payload, error) {
	payload := NewPayload(payloadData, tokenAliveDuration)
	token, err := builder.paseto.Encrypt(builder.symmetricKey, payload, nil)
	if err != nil {
		return "", nil, err
	}
	return token, payload, err
}
