// Command json performs similar functionality to:
// alias json='python -m json.tool'
// Unlike python 2, this script handles unicode nicely.
// In general, it adheres to the UNIX philosophy:
// Simply pipe unformatted JSON data to the command and formatted JSON comes out.
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
	fmt.Println(helps.FormatJSON(content))
}
