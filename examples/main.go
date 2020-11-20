package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/anshujalan/linkparser"
)

var htmlDoc = `
<html>
	<body>
		<h1>Hello World</h1>
		<a href="https://www.wikipedia.org">A link to wikipedia</a>
	</body>
</html>
`

func main() {
	r := strings.NewReader(htmlDoc)
	links, err := linkparser.Parse(r)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v", links)
}
