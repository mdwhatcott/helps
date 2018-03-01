package main

import (
	"crypto/sha256"

	"github.com/mdwhatcott/helps/nix"
)

func main() {
	nix.Hash(sha256.New())
}
