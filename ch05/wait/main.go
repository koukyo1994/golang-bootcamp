package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

// WaitForServerはURLのサーバへ接続を試みます
// 指数バックオフを使って1分間試みます。
// 全ての試みが失敗したらエラーを報告します。
func WaitForServer(url string) error {
	const timeout = 1 * time.Minute
	deadline := time.Now().Add(timeout)
	for tries := 0; time.Now().Before(deadline); tries++ {
		_, err := http.Head(url)
		if err == nil {
			return nil // 成功
		}
		log.Printf("server not responding (%s); retrying...", err)
		time.Sleep(time.Second << uint(tries)) // 指数バックオフ
	}
	return fmt.Errorf("server %s failed to respond after %s", url, timeout)
}

func main() {
	url := os.Args[1]
	if url == "" {
		log.Fatalf("Url not given. Aborting.\n")
	}
	err := WaitForServer(url)
	if err != nil {
		log.Fatalf("Site is down: %v\n", err)
	} else {
		log.Printf("Access to the url suceeded.\n")
	}
}
