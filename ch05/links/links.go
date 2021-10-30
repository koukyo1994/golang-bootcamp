package links

import (
	"fmt"
	"net/http"

	"golang.org/x/net/html"
)

// Extractは指定されたURLへHTTP GETリクエストを行い、レスポンスを
// HTMLとしてパースして、そのHTMLドキュメント内のリンクを返します。
func Extract(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("getting %s: %s", url, resp.Status)
	}

	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
	}

	var links []string
	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key != "href" {
					continue
				}
				link, err := resp.Request.URL.Parse(a.Val)
				if err != nil {
					continue // 不正なURLは無視
				}
				links = append(links, link.String())
			}
		}
	}
	forEachNode(doc, visitNode, nil)
	return links, nil
}

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
