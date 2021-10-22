package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func countElements(elementsCount map[string]int, n *html.Node) map[string]int {
	if n.Type == html.ElementNode {
		elementsCount[n.Data]++
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		elementsCount = countElements(elementsCount, c)
	}
	return elementsCount
}

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	elementsCount := make(map[string]int)
	for key, value := range countElements(elementsCount, doc) {
		fmt.Printf("%s: %d\n", key, value)
	}
}
