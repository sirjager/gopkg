package cache

import (
	"errors"
)

var (
	ErrNoRecord  = errors.New("record not found")
	ErrUnMarshal = errors.New("failed to unmarshal")
)
