package main

import (
	"bootcamp/ch05/links"
	"fmt"
	"log"
	"os"
)

// tokenは、20個の並行なリクエストという限界を
// 強制するために使われる計数セマフォです。
var tokens = make(chan struct{}, 20)

func crawl(url string) []string {
	fmt.Println(url)
	tokens <- struct{}{} // トークンを獲得
	list, err := links.Extract(url)
	<-tokens // トークンを解放
	if err != nil {
		log.Print(err)
	}
	return list
}

func main() {
	worklist := make(chan []string)
	var n int // worklistへの送信待ちの数

	// コマンドラインの引数で開始する
	n++
	go func() { worklist <- os.Args[1:] }()

	// ウェブを並行にクロールする
	seen := make(map[string]bool)
	for ; n > 0; n-- {
		list := <-worklist
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				n++
				go func(link string) {
					worklist <- crawl(link)
				}(link)
			}
		}
	}
}
