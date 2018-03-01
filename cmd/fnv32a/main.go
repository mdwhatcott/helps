package main

import (
	"hash/fnv"

	"github.com/mdwhatcott/helps/nix"
)

func main() {
	nix.Hash(fnv.New32a())
}
