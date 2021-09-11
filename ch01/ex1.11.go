package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}

	nbytes, err := io.Copy(io.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}

	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs  %7d  %s", secs, nbytes, url)
}

func main() {
	start := time.Now()
	ch := make(chan string)
	urls := []string{
		"https://www.google.com",
		"https://www.youtube.com",
		"https://www.yahoo.co.jp",
		"https://www.amazon.co.jp",
		"https://www.google.co.jp",
		"https://rakuten.co.jp",
		"https://www.wikipedia.org",
		"https://www.facebook.com",
		"https://zoom.us",
		"https://www.yahoo.com",
		"https://www.amazon.com",
		"https://twitter.com",
		"https://www.instagram.com",
		"https://www.tmall.com",
		"https://www.microsoft.com",
		"https://www.qq.com",
		"https://ameblo.jp",
		"https://recruit.japanpost.jp",
		"https://www.mercari.com/",
		"https://crowdworks.jp/",
		"https://www.google.com.sg/",
		"https://www.netflix.com/",
		"https://www.farfetch.com/jp/",
		"https://www.baidu.com/",
		"https://outlook.live.com",
		"https://fc2.com/",
		"https://www.sohu.com/",
		"https://www.office.com/",
		"https://deepl.com/",
		"https://world.taobao.com/",
		"https://lazada.sg/",
		"https://www.adobe.com/",
		"https://www.yoox.com/jp",
		"https://clickpost.jp/",
		"https://www.dmm.co.jp",
		"https://note.com/",
		"https://line.me/",
		"https://kakaku.com/",
		"https://www.asos.com/",
		"https://weblio.jp/",
		"https://jd.com/",
		"https://www.kuronekoyamato.co.jp/",
		"https://www.goo.ne.jp/",
		"https://www.dropbox.com/",
		"https://www.apple.com/",
		"https://matchfashion.com/",
		"https://rakuten-bank.co.jp/",
		"https://www.ebay.com/",
		"https://www.nhk.or.jp/",
		"https://www.360.cn/",
	}
	for _, url := range urls {
		go fetch(url, ch)
	}
	for range urls {
		fmt.Println(<-ch)
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}
