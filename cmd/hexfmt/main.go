// hexfmt takes a hex-formatted uint value (which Go outputs in %#v situations) and emits the base-10 value.
// $ echo '0x231c20139' | hexfmt // 9424732473
package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	all, err := io.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	raw := strings.TrimSpace(string(all))
	parsed, err := strconv.ParseUint(strings.TrimPrefix(raw, "0x"), 16, 64)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(parsed)
}
