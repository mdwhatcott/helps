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
	log.SetFlags(log.Llongfile)

	content, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	formatted, err := helps.FormatJSON(content)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(formatted))
}
