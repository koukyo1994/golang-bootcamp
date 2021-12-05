package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

// 送信用メッセージチャンネル
type client struct {
	ch   chan<- string
	name string
}

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string) // クライアントから受信する全てのメッセージ
)

func broadcaster() {
	clients := make(map[client]bool) // 全ての接続されているクライアント
	for {
		select {
		case msg := <-messages:
			// 受信するメッセージを全てのクライアントの
			// 送信用メッセージチャンネルへブロードキャストする
			for cli := range clients {
				cli.ch <- msg
			}
		case cli := <-entering:
			msg := "currently in the chat: "
			for cli := range clients {
				msg += cli.name + " "
			}
			cli.ch <- msg
			clients[cli] = true
		case cli := <-leaving:
			delete(clients, cli)
			close(cli.ch)
		}
	}
}

func handleConn(conn net.Conn) {
	ch := make(chan string) // 送信用のクライアントメッセージ
	go clientWriter(conn, ch)

	who := conn.RemoteAddr().String()
	ch <- "You are " + who
	messages <- who + " has arrived"
	entering <- client{ch, who}

	input := bufio.NewScanner(conn)
	for input.Scan() {
		messages <- who + ": " + input.Text()
	}
	// 注意: input.Err()からの潜在的なエラーを無視している

	leaving <- client{ch, who}
	messages <- who + " has left"
	conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg) // 注意: ネットワークのエラーを無視している
	}
}

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	go broadcaster()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}
