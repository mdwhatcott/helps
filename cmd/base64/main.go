package main

import (
	"encoding/base64"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	all, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	encoder := base64.NewEncoder(base64.StdEncoding, os.Stdout)
	_, err = encoder.Write(all)
	if err != nil {
		log.Fatal(err)
	}
	err = encoder.Close()
	if err != nil {
		log.Fatal(err)
	}
}
