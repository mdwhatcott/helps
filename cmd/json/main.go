// Command json performs similar functionality to:
// alias json='python -m json.tool'
// Unlike python 2, this script handles unicode nicely.
// In general, it adheres to the UNIX philosophy:
// Simply pipe unformatted JSON data to the command and formatted JSON comes out.
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/mdwhatcott/helps"
)

func main() {
	log.SetFlags(log.Lshortfile)

	if input, err := ioutil.ReadAll(os.Stdin); err != nil {
		log.Fatalf("Error reading from stdin: %s", err)
	} else if formatted, err := helps.FormatJSONSafe(input); err != nil {
		log.Fatalf("Error formatting JSON: %s\n%s", err, string(input))
	} else {
		fmt.Println(string(formatted))
	}
}
