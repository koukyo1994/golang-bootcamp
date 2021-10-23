package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

// CountWordsAndImagesはHTMLドキュメントに対するHTTP GETリクエストをurlへ行い、
// そのドキュメント内に含まれる単語と画像の数を返します。
func CountWordsAndImages(url string) (words, images int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		err = fmt.Errorf("parsing HTML: %s", err)
		return
	}
	words, images = countWordsAndImages(doc)
	return
}

func countWordsAndImages(n *html.Node) (words, images int) {
	images = countImages(n)
	text := strings.Join(getTexts(nil, n), " ")

	r := strings.NewReader(text)
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		words++
	}

	return
}

func getTexts(texts []string, n *html.Node) []string {
	if n.Type == html.TextNode {
		if n.Parent.Type != html.ElementNode || (n.Parent.Data != "style" && n.Parent.Data != "script") {
			texts = append(texts, n.Data)
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		texts = getTexts(texts, c)
	}
	return texts
}

func countImages(n *html.Node) (images int) {
	if n.Type == html.ElementNode && n.Data == "img" {
		images++
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		images += countImages(c)
	}
	return
}

func main() {
	for _, url := range os.Args[1:] {
		words, images, err := CountWordsAndImages(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "CountWordsAndImages failed with %s: %v\n", url, err)
			continue
		}
		fmt.Printf("%s:\n  words: %d\n  images %d\n", url, words, images)
	}
}
