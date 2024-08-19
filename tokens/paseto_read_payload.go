package tokens

func (b *pasetoBuilder) ReadPayload(payload *Payload, data interface{}) error {
	if err := payload.Valid(); err != nil {
		return err
	}
	if err := b.codec.Unmarshal(payload.Data, data); err != nil {
		return err
	}
	return nil
}
