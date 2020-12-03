package json

import (
	"encoding/json"

	"github.com/mailru/easyjson"
)

func Unmarshal(data []byte, v interface{}) error {
	if m, ok := v.(easyjson.Unmarshaler); ok {
		return easyjson.Unmarshal(data, m)
	}

	return json.Unmarshal(data, v)
}

func Deserialize(data string, v interface{}) error {
	return Unmarshal([]byte(data), v)
}
