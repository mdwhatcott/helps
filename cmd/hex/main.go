package main

import (
	"encoding/hex"
	"io"
	"log"
	"os"
)

func main() {
	all, err := io.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	encoder := hex.NewEncoder(os.Stdout)
	_, err = encoder.Write(all)
	if err != nil {
		log.Fatal(err)
	}
}
