package linkparser

import "io"

//Link refers to a link tag <a href=".."> in a HTML document
type Link struct {
	Href string
	Text string
}

//Parse parses a reader of a HTML file
func Parse(r io.Reader) ([]Link, error) {
	return nil, nil
}
