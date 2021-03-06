package main

import (
	"bootcamp/ch05/links"
	"fmt"
	"log"
	"os"
)

func crawl(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url)
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
			if !seen[link] {
				seen[link] = true
				unseenLinks <- link
			}
		}
	}
}
