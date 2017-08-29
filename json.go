package helps

import (
	"bytes"
	"encoding/json"
)

func FormatJSON(data []byte) (string, error) {
	var indent bytes.Buffer
	err := json.Indent(&indent, data, "", "  ")
	if err != nil {
		return "", err
	}
	return indent.String(), nil
}

func DumpJSON(v interface{}) string {
	b, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return err.Error()
	}

	return string(b)
}
