package helps

import "fmt"

func DumpFMT(v interface{}) string {
	return fmt.Sprintf("%#v", v)
}
