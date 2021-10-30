package main

import (
	"bootcamp/ch05/links"
	"fmt"
	"log"
	"os"
)

// breadthFirstはworklist内の個々の項目に対してfを呼び出します
// fから返された全ての項目はworlistへ追加されます
// fは、それぞれの項目に対して高々一度しか呼び出されません。
func breadthFirst(f func(item string) []string, worklist []string) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, f(item)...)
			}
		}
	}
}

func crawl(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}

func main() {
	// コマンドライン引数から開始し、
	// ウェブを幅優先探索でクロールする。
	breadthFirst(crawl, os.Args[1:])
}
