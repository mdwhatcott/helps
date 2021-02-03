// Command json performs similar functionality to:
// alias json='python -m json.tool'
// Unlike python 2, this script handles unicode nicely.
// In general, it adheres to the UNIX philosophy:
// Simply pipe JSON data to the command and formatted JSON comes out.
package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/mdwhatcott/helps"
)

func main() {
	log.SetFlags(log.Lshortfile)

	input, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatalf("Error reading from stdin: %s", err)
	}

	input = bytes.TrimSpace(input)
	if len(input) == 0 {
		return
	}

	formatted, err := helps.FormatJSONSafe(input)
	if err != nil {
		log.Fatalf("Error formatting JSON: %s\n%s", err, string(input))
	}

	fmt.Println(formatted)
}
