package helps

import (
	"bytes"
	"encoding/xml"
	"io"
)

// https://stackoverflow.com/a/27141132/605022
func FormatXML(data []byte) (string, error) {
	if len(data) > 3 && DumpFMT(data[:3]) == DumpFMT(utf8ByteOrderMark) {
		data = data[3:]
	}
	var buffer bytes.Buffer
	decoder := xml.NewDecoder(bytes.NewReader(data))
	encoder := xml.NewEncoder(&buffer)
	encoder.Indent("", "  ")
	for {
		token, err := decoder.Token()
		if err == io.EOF {
			encoder.Flush()
			return buffer.String(), nil
		}
		if err != nil {
			return "", err
		}
		err = encoder.EncodeToken(token)
		if err != nil {
			return "", err
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
