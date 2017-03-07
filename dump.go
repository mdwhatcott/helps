package helps

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func DumpJSON(v interface{}) string {
	b, err := json.Marshal(v)
	if err != nil {
		return err.Error()
	}

	var indent bytes.Buffer
	err = json.Indent(&indent, b, "", "  ")
	if err != nil {
		return err.Error()
	}
	return indent.String()
}

func DumpFMT(v interface{}) string {
	return fmt.Sprintf("%#v", v)
}
