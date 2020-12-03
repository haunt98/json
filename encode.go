package json

import (
	"encoding/json"

	"github.com/mailru/easyjson"
)

func Marshal(v interface{}) ([]byte, error) {
	if m, ok := v.(easyjson.Marshaler); ok {
		return easyjson.Marshal(m)
	}

	return json.Marshal(v)
}

func Serialize(v interface{}) (string, error) {
	data, err := Marshal(v)
	if err != nil {
		return "", err
	}

	return string(data), nil
}
