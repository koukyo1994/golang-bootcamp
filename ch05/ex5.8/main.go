package main

import (
	"fmt"
	"strings"

	"golang.org/x/net/html"
)

func ForEachNode(n *html.Node, pre, post func(n *html.Node) bool) bool {
	if pre != nil {
		if willContinue := pre(n); !willContinue {
			return willContinue
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if willContinue := ForEachNode(c, pre, post); !willContinue {
			return willContinue
		}
	}

	if post != nil {
		if willContinue := post(n); !willContinue {
			return willContinue
		}
	}
	return true
}

func hasId(n *html.Node, id string) bool {
	for _, a := range n.Attr {
		if a.Key == "id" && a.Val == id {
			return true
		}
	}
	return false
}

func ElementById(doc *html.Node, id string) *html.Node {
	var node *html.Node
	ForEachNode(doc, func(n *html.Node) bool {
		if n.Type == html.ElementNode && hasId(n, id) {
			node = n
			return false
		}
		return true
	}, nil)
	return node
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
			<p id="p"><a href="https://gopl.io">Gopl.io</a></p>
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
		</body>
	</html>`))
	n := ElementById(doc, "p")
	attrs := []string{}
	for _, a := range n.Attr {
		attrs = append(attrs, fmt.Sprintf("%s='%s'", a.Key, a.Val))
	}
	fmt.Printf("found node <%s %s>", n.Data, strings.Join(attrs, " "))
}
