package main

import (
	"crypto/sha512"
	"os"

	"github.com/mdwhatcott/helps/hashing"
)

func main() {
	hashing.Hash(os.Stdin, sha512.New(), os.Stdout)
}
