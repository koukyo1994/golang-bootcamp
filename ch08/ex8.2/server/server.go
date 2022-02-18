package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func main() {
	root, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatal(err)
	}

	log.Print("server started")
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		connection := Connection{
			CommandConn:      conn,
			DataConn:         nil,
			WorkingDirectory: Directory(root),
			RootDirectory:    Directory(root),
		}
		go handleConn(&connection)
	}
}

func handleConn(c *Connection) {
	_, err := io.WriteString(c.CommandConn, "connected\n")
	if err != nil {
		log.Print(err)
	}
	for {
		c.reply("ftp> ")
		cmd, args, err := c.readCommand()
		if err == io.EOF {
			break
		}

		switch cmd {
		case "ls":
			c.handlels(args)
		case "cd":
			c.handlecd(args)
		default:
			continue
		}
	}
}
