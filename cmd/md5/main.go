package main

import (
	"crypto/md5"
	"os"

	"github.com/mdwhatcott/helps/hashing"
)

func main() {
	hashing.Hash(os.Stdin, md5.New(), os.Stdout)
}
