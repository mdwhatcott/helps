package nix

import (
	"encoding/hex"
	"hash"
	"io"
	"log"
	"os"
)

func Hash(inner hash.Hash) {
	if _, err := io.Copy(inner, os.Stdin); err != nil {
		log.Fatal(err)
	}
	if _, err := hex.NewEncoder(os.Stdout).Write(inner.Sum(nil)); err != nil {
		log.Fatal(err)
	}
}
