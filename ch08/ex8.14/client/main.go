package main

import (
	"io"
	"log"
	"net"
	"os"
	"strings"
)

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}

func main() {
	name := os.Args[1]
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	mustCopy(conn, strings.NewReader(name))
	done := make(chan struct{})
	go func() {
		io.Copy(os.Stdout, conn)
		log.Println("done")
		done <- struct{}{} // メインgoroutineに通知
	}()
	mustCopy(conn, os.Stdin)
	conn.Close()
	<-done // バックグラウンドのゴルーチンが完了するのを待つ
}
