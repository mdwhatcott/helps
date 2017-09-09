package helps

import (
	"bytes"
	"encoding/json"
	"log"
)

func FormatJSON(data []byte) string {
	formatted, err := FormatJSONSafe(data)
	if err != nil {
		log.Panic(err)
	}
	return formatted
}

func FormatJSONSafe(data []byte) (string, error) {
	var indent bytes.Buffer
	err := json.Indent(&indent, data, "", "  ")
	if err != nil {
		return "", err
	}
	return indent.String(), nil
}

func DumpJSON(v interface{}) string {
	dump, err := DumpJSONSafe(v)
	if err != nil {
		log.Panic(err)
	}
	return dump
}

func DumpJSONSafe(v interface{}) (string, error) {
	b, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return "", err
	}

	return string(b), nil
}
