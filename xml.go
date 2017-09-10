package helps

import (
	"bytes"
	"encoding/xml"
	"io"
	"log"
)

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

func FormatXML(data []byte) string {
	formatted, err := FormatXMLSafe(data)
	if err != nil {
		log.Panic(err)
	}
	return formatted
}

func FormatXMLSafe(data []byte) (string, error) {
	return NewXMLFormatter(data).Format()
}

type XMLFormatter struct {
	buffer  *bytes.Buffer
	encoder *xml.Encoder
	decoder *xml.Decoder
}

func NewXMLFormatter(data []byte) *XMLFormatter {
	buffer := new(bytes.Buffer)
	return &XMLFormatter{
		buffer:  buffer,
		encoder: setupEncoder(buffer),
		decoder: setupDecoder(data),
	}
}

// https://stackoverflow.com/a/27141132/605022
func (this *XMLFormatter) Format() (string, error) {
	for {
		token, err := this.decoder.Token()
		if err == io.EOF {
			return this.finalResult()
		} else if err != nil {
			return "", err
		}

		err = this.encoder.EncodeToken(token)
		if err != nil {
			return "", err
		}
	}
}

func (this *XMLFormatter) finalResult() (string, error) {
	this.encoder.Flush()
	return this.buffer.String(), nil
}

func setupDecoder(data []byte) *xml.Decoder {
	return xml.NewDecoder(bytes.NewReader(trimByteOrderMark(data)))
}

func setupEncoder(buffer io.Writer) *xml.Encoder {
	encoder := xml.NewEncoder(buffer)
	encoder.Indent("", "  ")
	return encoder
}

func trimByteOrderMark(data []byte) []byte {
	return bytes.TrimLeft(data, string(utf8ByteOrderMark))
}

var utf8ByteOrderMark = []byte{239, 187, 191}
