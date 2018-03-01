package main

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"log"
	"os"
)

func main() {
	hasher := md5.New()
	_, err := io.Copy(hasher, os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	encoder := hex.NewEncoder(os.Stdout)
	_, err = encoder.Write(hasher.Sum(nil))
	if err != nil {
		log.Fatal(err)
	}
}
