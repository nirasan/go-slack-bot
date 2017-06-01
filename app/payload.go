package app

import (
	"encoding/json"
	"io"
)

type Payload map[string]interface{}

func DecodeJSON(r io.Reader) (Payload, error) {
	data := make(map[string]interface{})
	if err := json.NewDecoder(r).Decode(&data); err != nil {
		return nil, err
	}
	return data, nil
}

func (p Payload) String(key string) string {
	if v, ok := p[key]; !ok {
		return ""
	} else if vv, ok := v.(string); !ok {
		return ""
	} else {
		return vv
	}
}

func (p Payload) Type() string {
	return p.String("type")
}
