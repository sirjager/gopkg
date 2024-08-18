package serializer

import (
	"github.com/vmihailenco/msgpack/v5"
)

type msgpackSerializer struct{}

func NewMsgPackSerializer() Serializer {
	return &msgpackSerializer{}
}

// Marshal usage:- data, err := msgpack.Marshal(&user)
func (m *msgpackSerializer) Marshal(v interface{}) ([]byte, error) {
	return msgpack.Marshal(v)
}

// Unmarshal usage:- err := msgpack.Unmarshal(data, &value)
func (m *msgpackSerializer) Unmarshal(data []byte, v interface{}) error {
	return msgpack.Unmarshal(data, v)
}
