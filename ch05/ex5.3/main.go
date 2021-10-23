package main

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func visit(links []string, n *html.Node) []string {
	if n.Type == html.TextNode {
		if n.Parent.Type != html.ElementNode || (n.Parent.Data != "style" && n.Parent.Data != "script") {
			links = append(links, n.Data)
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}
	return links
}

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ex5.3: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("%s\n", strings.Join(visit(nil, doc), ""))
}
