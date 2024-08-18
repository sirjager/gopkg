package tokens

// VerifyToken verifies a token
func (b *pasetoBuilder) VerifyToken(token string, data interface{}) (*Payload, error) {
	payload := &Payload{}
	err := b.paseto.Decrypt(token, b.symmetricKey, payload, nil)
	if err != nil {
		return nil, ErrInvalidToken
	}
	if err = payload.Valid(); err != nil {
		return nil, err
	}
	if err = b.codec.Unmarshal(payload.Data, data); err != nil {
		return nil, err
	}
	return payload, nil
}
