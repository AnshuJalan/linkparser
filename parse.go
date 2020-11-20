package linkparser

import (
	"io"
	"log"

	"golang.org/x/net/html"
)

//Link refers to a link tag <a href=".."> in a HTML document
type Link struct {
	Href string
	Text string
}

//Parse parses a reader of a HTML file
func Parse(r io.Reader) ([]Link, error) {
	root, err := html.Parse(r)
	if err != nil {
		log.Fatal(err)
	}
	linkNodes := findLinkNodes(root)
	links := buildLinks(linkNodes)

	return links, nil
}

func buildLinks(linkNodes []*html.Node) []Link {
	var links []Link
	for _, node := range linkNodes {
		for _, a := range node.Attr {
			if a.Key == "href" {
				links = append(links, Link{a.Val, node.FirstChild.Data})
				break
			}
		}
	}
	return links
}

func findLinkNodes(node *html.Node) []*html.Node {
	if node.Type == html.ElementNode && node.Data == "a" {
		return []*html.Node{node}
	}
	var res []*html.Node
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		res = append(res, findLinkNodes(c)...)
	}

	return res
}
