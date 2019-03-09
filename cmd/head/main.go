package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

var (
	n int
	c int
)

func main() {
	flag.IntVar(&n, "n", 10, "lines")
	flag.IntVar(&c, "c", 0, "bytes")
	flag.Parse()

	if n == 0 && c == 0 {
		log.Fatal("Illegal line count:", n)
	}

	args := flag.Args()
	var in io.Reader = os.Stdin
	if len(args) > 0 {
		file, err := os.Open(args[len(args)-1])
		if err != nil {
			log.Fatal(err)
		}
		defer func() {
			err = file.Close()
			if err != nil {
				log.Fatal(err)
			}
		}()
		in = file
	}

	if c > 0 {
		buffer := make([]byte, c)
		in.Read(buffer)
	}
	scanner := bufio.NewScanner(in)
	for ; scanner.Scan() && n >= 0; n-- {
		fmt.Println(scanner.Text())
	}
}
