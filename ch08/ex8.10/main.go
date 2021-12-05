package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

var done = make(chan struct{})

func cancelled() bool {
	select {
	case <-done:
		return true
	default:
		return false
	}
}

// Extractは指定されたURLへHTTP GETリクエストを行い、レスポンスを
// HTMLとしてパースして、そのHTMLドキュメント内のリンクを返します。
func Extract(url string) ([]string, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	// 今だとNewRequestWithContextを使うのが一般的っぽい？
	req.Cancel = done
	resp, err := http.DefaultClient.Do(req)
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

func crawl(url string) []string {
	fmt.Println(url)
	list, err := Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}

func main() {
	worklist := make(chan []string)  // URLのリスト、重複を含む
	unseenLinks := make(chan string) // 重複していないURL

	// コマンドライン引数をworklistへ追加する
	go func() { worklist <- os.Args[1:] }()

	// 入力を検出するとキャンセルできる
	go func() {
		os.Stdin.Read(make([]byte, 1))
		close(done)
	}()

	// 未探索のリンクを取得するために20個のクローラのゴルーチンを生成する。
	for i := 0; i < 20; i++ {
		go func() {
			for link := range unseenLinks {
				foundLinks := crawl(link)
				go func() { worklist <- foundLinks }()
			}
		}()
	}

	// メインゴルーチンはworklistの項目の重複をなくし、
	// 未探索の項目をクローラに送る
	seen := make(map[string]bool)
	for list := range worklist {
		for _, link := range list {
			if cancelled() {
				return
			}
			if !seen[link] {
				seen[link] = true
				unseenLinks <- link
			}
		}
	}
}
