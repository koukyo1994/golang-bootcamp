package main

import (
	"fmt"
	"strings"

	"golang.org/x/net/html"
)

func ElementsByTagName(doc *html.Node, name ...string) []*html.Node {
	nodes := []*html.Node{}
	var gatherNode func(n *html.Node, tags []string)
	gatherNode = func(n *html.Node, tags []string) {
		if isin(n.Data, tags) {
			nodes = append(nodes, n)
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			gatherNode(c, tags)
		}
	}

	gatherNode(doc, name)
	return nodes
}

func isin(tag string, tags []string) bool {
	for _, t := range tags {
		if t == tag {
			return true
		}
	}
	return false
}

func main() {
	doc, _ := html.Parse(strings.NewReader(`
	<html lang="en">
		<head>
			<title>image</title>
		</head>
		<body>
			<!-- コメント -->
			<h1>リンクのサンプル</h1>
			<p><a href="https://gopl.io">Gopl.io</a></p>
			<h1>画像のサンプル</h1>
			<img src='image.png'>
			<h2>テーブルのサンプル</h2>
			<table>
				<tr style='text-align: left'>
					<th>Item</th>
					<th>Price</th>
				</tr>
				<tr>
					<td>aaa</td>
					<td>$10</td>
				</tr>
			</table>
			<img src='image2.png'>
		</body>
	</html>`))
	images := ElementsByTagName(doc, "img")
	fmt.Printf("%v\n", images)

	headings := ElementsByTagName(doc, "h1", "h2", "h3", "h4")
	fmt.Printf("%v\n", headings)
}
