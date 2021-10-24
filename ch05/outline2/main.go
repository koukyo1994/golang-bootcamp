package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

var depth int

// forEachNodeはnから始まるツリー内の個々のノードxに対して
// 関数pre(x)とpost(x)を呼び出します。その二つの関数はオプションです。
// preは子ノードを訪れる前に呼び出されます。(前順:preorder)
// postは子ノードを訪れた後に呼び出されます。(後順:postorder)
func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}

	if post != nil {
		post(n)
	}
}

func startElement(n *html.Node) {
	if n.Type == html.ElementNode {
		fmt.Printf("%*s<%s>\n", depth*2, "", n.Data)
		depth++
	}
}

func endElement(n *html.Node) {
	if n.Type == html.ElementNode {
		depth--
		fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
	}
}

func main() {
	url := os.Args[1]
	if url == "" {
		fmt.Fprintf(os.Stderr, "Url is not given\n")
		os.Exit(1)
	}
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "reading %s: %v\n", url, err)
		os.Exit(1)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		fmt.Fprintf(os.Stderr, "reading %s: got http status %d\n", url, resp.StatusCode)
		os.Exit(1)
	}
	doc, err := html.Parse(resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "parsing HTML from %s: %v\n", url, err)
		os.Exit(1)
	}
	forEachNode(doc, startElement, endElement)
}
