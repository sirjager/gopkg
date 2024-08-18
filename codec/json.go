package codec

import "encoding/json"

type _json struct{}

func NewJSON() Codec {
	return &_json{}
}

func (j *_json) Marshal(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

func (j *_json) Unmarshal(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}
