package tokens

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/sirjager/gopkg/utils"
)

func TestVerifyToken(t *testing.T) {
	payloadData := &testPayloadData{
		UserID: utils.RandomString(32),
	}
	token, validPayload := testCreateToken(t, payloadData)
	expiredToken, _ := testCreateToken(t, payloadData, time.Millisecond)
	time.Sleep(time.Second) // Waiting a second to expire our token

	testCases := []struct {
		test  func(payload *Payload, err error)
		name  string
		token string
	}{
		{
			name:  "VerifyToken/Valid",
			token: token,
			test: func(payload *Payload, err error) {
				require.NoError(t, err)
				require.NotEmpty(t, payload)
				data, ok := validPayload.Payload.(*testPayloadData)
				require.True(t, ok)
				require.Equal(t, payloadData.UserID, data.UserID)
				require.Equal(t, payloadData.UserID, data.UserID)
			},
		},
		{
			name:  "VerifyToken/Invalid",
			token: utils.RandomString(64),
			test: func(payload *Payload, err error) {
				require.Error(t, err)
				require.Empty(t, payload)
				require.Equal(t, err.Error(), ErrInvalidToken.Error())
				require.ErrorIs(t, err, ErrInvalidToken)
			},
		},
		{
			name:  "VerifyToken/Expired",
			token: expiredToken,
			test: func(payload *Payload, err error) {
				require.Error(t, err)
				require.Empty(t, payload)
				require.Equal(t, err.Error(), ErrExpiredToken.Error())
				require.ErrorIs(t, err, ErrExpiredToken)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.test(pasetoTokenBuilder.VerifyToken(tc.token))
		})
	}
}
