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

	content, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	formatted, err := helps.FormatXML(content)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(formatted))
}
