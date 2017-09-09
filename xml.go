package helps

import (
	"bytes"
	"encoding/xml"
	"io"
	"log"
)

func FormatXML(data []byte) string {
	formatted, err := FormatXMLSafe(data)
	if err != nil {
		log.Panic(err)
	}
	return formatted
}

// https://stackoverflow.com/a/27141132/605022
func FormatXMLSafe(data []byte) (string, error) {
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
	dump, err := DumpXMLSafe(v)
	if err != nil {
		log.Panic(err)
	}
	return dump
}

func DumpXMLSafe(v interface{}) (string, error) {
	dump, err := xml.MarshalIndent(v, "", "  ")
	if err != nil {
		return "", err
	}

	return string(dump), nil
}
