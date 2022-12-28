package hashing

import (
	"encoding/hex"
	"hash"
	"io"
	"log"
)

func Hash(reader io.Reader, hasher hash.Hash, writer io.Writer) {
	_, err := io.Copy(hasher, reader)
	if err != nil {
		log.Fatal(err)
	}
	encoder := hex.NewEncoder(writer)
	_, err = encoder.Write(hasher.Sum(nil))
	if err != nil {
		log.Fatal(err)
	}
}
