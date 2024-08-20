package tokens

import (
	"fmt"

	"github.com/aead/chacha20poly1305"
	"github.com/o1egl/paseto"

	"github.com/sirjager/gopkg/codec"
)

const PasetoSymmetricKeyLength = chacha20poly1305.KeySize

// PasetoBuilder implements TokenBuilder
type PasetoBuilder struct {
	codec        codec.Codec
	paseto       *paseto.V2
	symmetricKey []byte
}

// NewPasetoBuilder returns a new PasetoBuilder
func NewPasetoBuilder(symmetricKey string, codecs ...codec.Codec) (TokenBuilder, error) {
	if len(symmetricKey) != PasetoSymmetricKeyLength {
		return nil, fmt.Errorf(
			"invalid key size: must be exactly %d characters",
			PasetoSymmetricKeyLength,
		)
	}

	var selected codec.Codec
	if len(codecs) == 1 && codecs[0] != nil {
		selected = codecs[0]
	} else {
		selected = codec.NewJSON()
	}

	builder := &PasetoBuilder{
		paseto:       paseto.NewV2(),
		symmetricKey: []byte(symmetricKey),
		codec:        selected,
	}

	return builder, nil
}
