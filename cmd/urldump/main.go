package main

import (
	"flag"
	"fmt"
	"log"
	"net/url"
	"strings"
)

func main() {
	flag.Usage = func() {
		log.SetFlags(0)
		log.Println("Feed space-separated args (ex. a full curl command) " +
			"to this program and it will parse each as a url, printing " +
			"any that have a non-blank hostname and/or at least one " +
			"slash in their path.")
	}
	flag.Parse()
	for _, arg := range flag.Args() {
		parsed, err := url.Parse(arg)
		if err == nil && (len(parsed.Host) > 0 || strings.Contains(parsed.Path, "/")) {
			Print("Scheme", parsed.Scheme)
			Print("User", parsed.User.String())
			Print("Host", parsed.Host)
			Print("Path", parsed.Path)
			Print("Query", parsed.RawQuery)
			Print("Fragment", parsed.Fragment)
			query := parsed.Query()
			for key := range query {
				Print("  "+key, query.Get(key))
			}
		}
	}
}

func Print(title, value string) {
	if len(value) > 0 {
		fmt.Printf(title+":\t[%s]\n", value)
	}
}
