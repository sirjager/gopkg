package tokens

import (
	"time"
)

// CreateToken creates a new token
func (b *pasetoBuilder) CreateToken(data interface{}, exp time.Duration) (string, *Payload, error) {
	payload, err := newPayload(data, exp, b.codec)
	if err != nil {
		return "", nil, err
	}
	token, err := b.paseto.Encrypt(b.symmetricKey, payload, nil)
	if err != nil {
		return "", nil, err
	}
	return token, payload, err
}
