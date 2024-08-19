package tokens

// VerifyToken verifies a token
func (b *pasetoBuilder) VerifyToken(token string, data interface{}) (*Payload, error) {
	payload := &Payload{}
	err := b.paseto.Decrypt(token, b.symmetricKey, payload, nil)
	if err != nil {
		return nil, ErrInvalidToken
	}
	if err := b.ReadPayload(payload, data); err != nil {
		return nil, err
	}
	return payload, nil
}
