package serializer

import "encoding/json"

type jsonSerializer struct{}

func NewJSONSerializer() *jsonSerializer {
	return &jsonSerializer{}
}

func (j *jsonSerializer) Marshal(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

func (j *jsonSerializer) Unmarshal(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}
