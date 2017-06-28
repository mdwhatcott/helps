package helps

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func FormatJSON(data []byte) string {
	var indent bytes.Buffer
	err := json.Indent(&indent, data, "", "  ")
	if err != nil {
		return err.Error()
	}
	return indent.String()
}

func DumpJSON(v interface{}) string {
	b, err := json.Marshal(v)
	if err != nil {
		return err.Error()
	}

	return FormatJSON(b)
}

func DumpFMT(v interface{}) string {
	return fmt.Sprintf("%#v", v)
}
