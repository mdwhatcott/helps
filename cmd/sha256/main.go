package main

import (
	"crypto/sha256"
	"os"

	"github.com/mdwhatcott/helps/hashing"
)

func main() {
	hashing.Hash(os.Stdin, sha256.New(), os.Stdout)
}
