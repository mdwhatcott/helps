package main

import (
	"crypto/sha1"
	"os"

	"github.com/mdwhatcott/helps/hashing"
)

func main() {
	hashing.Hash(os.Stdin, sha1.New(), os.Stdout)
}
