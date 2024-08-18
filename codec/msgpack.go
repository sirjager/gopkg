package codec

import (
	"github.com/vmihailenco/msgpack/v5"
)

type _msgpack struct{}

func NewMsgPack() Codec {
	return &_msgpack{}
}

// Marshal usage:- data, err := msgpack.Marshal(&user)
func (m *_msgpack) Marshal(v interface{}) ([]byte, error) {
	return msgpack.Marshal(v)
}

// Unmarshal usage:- err := msgpack.Unmarshal(data, &value)
func (m *_msgpack) Unmarshal(data []byte, v interface{}) error {
	return msgpack.Unmarshal(data, v)
}
