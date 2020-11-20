package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/anshujalan/linkparser"
)

func main() {
	filename := flag.String("file", "index.html", "Html to be parsed.")
	flag.Parse()

	file, err := os.Open(*filename)
	if err != nil {
		log.Fatal(err)
	}

	links, err := linkparser.Parse(file)
	if err != nil {
		log.Fatal(err)
	}

	for _, link := range links {
		fmt.Printf("%+v\n\n", link)
	}
}
