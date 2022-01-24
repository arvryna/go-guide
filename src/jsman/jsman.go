package jsman

// This class can pack and unpack JSON, but jsonstr and struct of JSON must be provided
// in advance

import (
	"encoding/json"
)

type JsonParser struct {
	JsonStr        string
	packedData     []byte
	useDefaultData bool
}

func (t *JsonParser) Pack() {
}

// Unpack the json string into "Any" interface
func (t *JsonParser) Unpack(i interface{}) error {
	// Go's marshal library requires jsonstr as byte slice
	data := []byte(t.JsonStr)

	err := json.Unmarshal(data, i)
	if err != nil {
		return err
	}

	return nil
}

// If you use type of pointer like below
func (t *JsonParser) UpdateJSONstr(jsonStr string) {
	t.JsonStr = jsonStr
}

// *** Usage sample ****
