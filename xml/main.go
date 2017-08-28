// Command XML indents XML from stdin to stdout
// In general, it adheres to the UNIX philosophy:
// Simply pipe unformatted XML data to the command and formatted XML comes out.
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/mdwhatcott/helps"
)

func main() {
	log.SetFlags(log.Llongfile)

	if input, err := ioutil.ReadAll(os.Stdin); err != nil {
		log.Fatalf("Error reading from stdin: %s", err)
	} else if formatted, err := helps.FormatXML(input); err != nil {
		log.Fatalf("Error formatting XML: %s\n%s", err, string(input))
	} else {
		fmt.Println(string(formatted))
	}
}
