package tokens

import (
	"fmt"

	"github.com/aead/chacha20poly1305"
	"github.com/o1egl/paseto"
)

const PasetoSymmetricKeyLength = chacha20poly1305.KeySize

// PasetoBuilder implements TokenBuilder
type PasetoBuilder struct {
	paseto       *paseto.V2
	symmetricKey []byte
}

// NewPasetoBuilder returns a new PasetoBuilder
func NewPasetoBuilder(symmetricKey string) (TokenBuilder, error) {
	if len(symmetricKey) != PasetoSymmetricKeyLength {
		return nil, fmt.Errorf(
			"invalid key size: must be exactly %d characters",
			PasetoSymmetricKeyLength,
		)
	}
	builder := &PasetoBuilder{
		paseto:       paseto.NewV2(),
		symmetricKey: []byte(symmetricKey),
	}

	return builder, nil
}
