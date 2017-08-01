// Command XML indents XML from stdin to stdout
// In general, it adheres to the UNIX philosophy:
// Simply pipe unformatted XML data to the command and formatted XML comes out.
package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/mdwhatcott/helps"
)

func main() {
	content, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(string(content))
	// TODO: non-zero return code when the input isn't valid XML
	fmt.Println(helps.FormatXML(content))
}
