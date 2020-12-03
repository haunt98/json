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
