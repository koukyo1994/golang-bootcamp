package main

import (
	"bootcamp/ch05/links"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

const ROOT = "./assets/ex5.13-content/"

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

func crawl(u string) []string {
	fmt.Println(u)
	list, err := links.Extract(u)
	if err != nil {
		log.Print(err)
	}
	result, err := url.Parse(u)
	if err != nil {
		log.Print(err)
	}
	domain := strings.TrimPrefix(result.Host, "www.")
	err = saveContent(u)
	if err != nil {
		log.Print(err)
	}
	inDomainList := []string{}
	for _, link := range list {
		res, err := url.Parse(link)
		if err != nil {
			log.Print(err)
		}
		if strings.TrimPrefix(res.Host, "www.") == domain {
			inDomainList = append(inDomainList, link)
		}
	}
	return inDomainList
}

func saveContent(u string) error {
	result, err := url.Parse(u)
	if err != nil {
		return err
	}
	domain := strings.TrimPrefix(result.Host, "www.")
	path := ROOT + domain + result.Path
	if strings.HasSuffix(path, "/") || result.Path == "" {
		if os.MkdirAll(path, 0777) != nil {
			return err
		}
	} else if strings.Contains(result.Path, "/") {
		splittedPath := strings.Split(result.Path, "/")
		if os.MkdirAll(ROOT+domain+"/"+strings.Join(splittedPath[:len(splittedPath)-1], "/"), 0777) != nil {
			return err
		}
	}
	if strings.HasSuffix(path, "/") {
		path += "index.html"
	} else if result.Path == "" {
		path += "/index.html"
	}
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	resp, err := http.Get(u)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("getting %s: %s", u, resp.Status)
	}
	_, err = io.Copy(f, resp.Body)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	// スクレイピングするURLを指定します
	// domains := os.Args[1]
	// コマンドライン引数から開始し、
	// ウェブを幅優先探索でクロールする。
	breadthFirst(crawl, os.Args[1:])
}
