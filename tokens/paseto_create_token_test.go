package tokens

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/sirjager/gopkg/utils"
)

type testPayloadData struct {
	UserID string `json:"user_id,omitempty"`
}

func testCreateToken(t *testing.T, payloadData *testPayloadData, duration ...time.Duration) (string, *Payload) {
	tokenAliveDuration := time.Second * 10
	if len(duration) == 1 {
		tokenAliveDuration = duration[0]
	}
	token, payload, err := pasetoTokenBuilder.CreateToken(payloadData, tokenAliveDuration)
	require.NoError(t, err)
	require.NotEmpty(t, token)
	require.NotEmpty(t, payload)
	return token, payload
}

func TestCreateToken(t *testing.T) {
	testCases := []struct {
		test        func(token string, payload *Payload, err error)
		payloadData *testPayloadData
		name        string
	}{
		{
			name: "CreateToken/OK",
			payloadData: &testPayloadData{
				UserID: utils.RandomString(32),
			},
			test: func(token string, payload *Payload, err error) {
				require.NoError(t, err)
				require.NotEmpty(t, token)
				require.NotEmpty(t, payload)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.test(pasetoTokenBuilder.CreateToken(tc.payloadData, time.Minute))
		})
	}
}
