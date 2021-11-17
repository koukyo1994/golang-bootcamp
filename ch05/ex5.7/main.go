package main

import (
	"fmt"
	"strings"

	"golang.org/x/net/html"
)

var depth int

// ForEachNodeはnから始まるツリー内の個々のノードxに対して
// 関数pre(x)とpost(x)を呼び出します。その二つの関数はオプションです。
// preは子ノードを訪れる前に呼び出されます。(前順:preorder)
// postは子ノードを訪れた後に呼び出されます。(後順:postorder)
func ForEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ForEachNode(c, pre, post)
	}

	if post != nil {
		post(n)
	}
}

func StartElement(n *html.Node) {
	switch n.Type {
	case html.CommentNode:
		fmt.Printf("%*s<!--%s-->\n", depth*2, "", n.Data)
	case html.ElementNode:
		output := fmt.Sprintf("%*s<%s", depth*2, "", n.Data)
		for _, a := range n.Attr {
			output += fmt.Sprintf(" %s='%s'", a.Key, a.Val)
		}
		if n.FirstChild == nil {
			output += "/>"
		} else {
			output += ">"
			depth++
		}
		fmt.Printf("%s\n", output)
	case html.TextNode:
		if strings.Contains(n.Data, "\n") {
			return
		}
		s := strings.TrimSpace(n.Data)
		fmt.Printf("%*s%s\n", depth*2, "", s)
	}
}

func EndElement(n *html.Node) {
	if n.Type == html.ElementNode && n.FirstChild != nil {
		depth--
		fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
	}
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
		</body>
	</html>`))
	ForEachNode(doc, StartElement, EndElement)
}
