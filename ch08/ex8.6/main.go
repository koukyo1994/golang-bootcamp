package main

import (
	"bootcamp/ch05/links"
	"flag"
	"fmt"
	"log"
)

var depth = flag.Int("depth", 1, "depth limit to extract links from given link")

// tokenは、20個の並行なリクエストという限界を
// 強制するために使われる計数セマフォです。
var tokens = make(chan struct{}, 20)

type linkWithDepth struct {
	link  string
	depth int
}

func crawl(url string, depth int) []linkWithDepth {
	fmt.Printf("depth: %d, link: %s\n", depth, url)
	tokens <- struct{}{}
	list, err := links.Extract(url)
	<-tokens
	if err != nil {
		log.Print(err)
	}
	linksWithDepth := make([]linkWithDepth, 0)
	for _, link := range list {
		linksWithDepth = append(linksWithDepth, linkWithDepth{link: link, depth: depth + 1})
	}
	return linksWithDepth
}

func main() {
	flag.Parse()

	worklist := make(chan []linkWithDepth)
	var n int // worklistへの送信待ちの数

	// コマンドラインの引数で開始する
	n++
	argsWithDepth := make([]linkWithDepth, 0)
	for _, arg := range flag.Args() {
		argsWithDepth = append(argsWithDepth, linkWithDepth{arg, 0})
	}
	go func() { worklist <- argsWithDepth }()

	// ウェブを並行にクロールする
	seen := make(map[string]bool)
	for ; n > 0; n-- {
		list := <-worklist
		for _, link := range list {
			if !seen[link.link] && (link.depth <= *depth) {
				seen[link.link] = true
				n++
				go func(link string, depth int) {
					worklist <- crawl(link, depth)
				}(link.link, link.depth)
			}
		}
	}
}
