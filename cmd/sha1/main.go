package main

import (
	"crypto/sha1"

	"github.com/mdwhatcott/helps/nix"
)

func main() {
	nix.Hash(sha1.New())
}
