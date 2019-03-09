package replace

import "strings"

type Replacer func() (OLD, NEW string)

func All(text string, replacements ...Replacer) string {
	for _, r := range replacements {
		o, n := r()
		text = strings.ReplaceAll(text, o, n)
	}
	return text
}

func Ment(old, new string) Replacer {
	return func() (string, string) { return old, new }
}
