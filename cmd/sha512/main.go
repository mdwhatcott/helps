package main

import (
	"crypto/sha512"

	"github.com/mdwhatcott/helps/nix"
)

func main() {
	nix.Hash(sha512.New())
}
