package main

import (
	"encoding/hex"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	all, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	encoder := hex.NewEncoder(os.Stdout)
	_, err = encoder.Write(all)
	if err != nil {
		log.Fatal(err)
	}
}
