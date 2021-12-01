package main

import (
	"flag"
	"io"
	"log"
	"net"
	"strconv"
	"time"
)

var port = flag.Int("port", 0, "port number for the server to listen")

func main() {
	flag.Parse()
	localhost := "localhost:" + strconv.Itoa(*port)

	listener, err := net.Listen("tcp", localhost)
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) // 接続が切れた
			continue
		}
		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()
	for {
		_, err := io.WriteString(conn, time.Now().Format("15:04:05\n"))
		if err != nil {
			return // 例: クライアントとの接続が切れた
		}
		time.Sleep(1 * time.Second)
	}
}
