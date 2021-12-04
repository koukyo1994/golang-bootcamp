package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

func handleConn(c net.Conn) {
	input := bufio.NewScanner(c)
	ticker := time.NewTicker(1 * time.Second)
	reset := make(chan struct{})
	count := 10
	go func() {
		for {
			select {
			case <-ticker.C:
				count--
			case <-reset:
				count = 10
			}
		}
	}()

	go func() {
		for {
			if count < 0 {
				c.Close()
				ticker.Stop()
				return
			}
		}
	}()
	for input.Scan() {
		go echo(c, input.Text(), 1*time.Second)
		reset <- struct{}{}
	}
	c.Close()
	ticker.Stop()
}

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
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
