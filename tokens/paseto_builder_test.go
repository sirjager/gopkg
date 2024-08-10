package tokens

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/sirjager/gopkg/utils"
)

var pasetoTokenBuilder TokenBuilder

func TestNewPasetoBuilder(t *testing.T) {
	testCases := []struct {
		test  func(builder TokenBuilder, err error)
		name      string
		secretKey string
	}{
		{
			name:      "SmallSecretKey",
			secretKey: utils.RandomString(PasetoSymmetricKeyLength - 2),
			test: func(builder TokenBuilder, err error) {
				require.Error(t, err)
				require.Empty(t, builder)
			},
		},
		{
			name:      "LargeSecretKey",
			secretKey: utils.RandomString(PasetoSymmetricKeyLength + 2),
			test: func(builder TokenBuilder, err error) {
				require.Error(t, err)
				require.Empty(t, builder)
			},
		},
		{
			name:      "ValidSecretKey",
			secretKey: utils.RandomString(PasetoSymmetricKeyLength),
			test: func(builder TokenBuilder, err error) {
				require.NoError(t, err)
				require.NotEmpty(t, builder)
				pasetoTokenBuilder = builder
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.test(NewPasetoBuilder(tc.secretKey))
		})
	}
}
