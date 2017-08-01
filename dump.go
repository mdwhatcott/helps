package helps

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
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
	b, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return err.Error()
	}

	return string(b)
}

// https://stackoverflow.com/a/27141132/605022
func FormatXML(data []byte) string {
	if len(data) > 3 && DumpFMT(data[:3]) == DumpFMT(utf8ByteOrderMark) {
		data = data[3:]
	}
	b := &bytes.Buffer{}
	decoder := xml.NewDecoder(bytes.NewReader(data))
	encoder := xml.NewEncoder(b)
	encoder.Indent("", "  ")
	for {
		token, err := decoder.Token()
		if err == io.EOF {
			encoder.Flush()
			return b.String()
		}
		if err != nil {
			return err.Error()
		}
		err = encoder.EncodeToken(token)
		if err != nil {
			return err.Error()
		}
	}
}

var utf8ByteOrderMark = []byte{239, 187, 191}

func DumpXML(v interface{}) string {
	b, err := xml.MarshalIndent(v, "", "  ")
	if err != nil {
		return err.Error()
	}

	return string(b)
}

func DumpFMT(v interface{}) string {
	return fmt.Sprintf("%#v", v)
}
