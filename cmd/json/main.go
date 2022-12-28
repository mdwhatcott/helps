package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	input, err := io.ReadAll(os.Stdin)
	if err != nil {
		log.Fatalf("Error reading from stdin: %s", err)
	}

	input = bytes.TrimSpace(input)
	if len(input) == 0 {
		return
	}

	var indent bytes.Buffer
	err = json.Indent(&indent, input, "", "  ")
	if err != nil {
		log.Fatalf("Error formatting JSON: %s\n%s", err, string(input))
	}
	fmt.Println(indent.String())
}
