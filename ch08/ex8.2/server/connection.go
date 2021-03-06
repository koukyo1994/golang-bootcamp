package main

import (
	"bufio"
	"net"
	"strings"
)

type Connection struct {
	CommandConn      net.Conn
	DataConn         net.Conn
	WorkingDirectory Directory
	RootDirectory    Directory
}

func (c *Connection) readCommand() (command string, args []string, err error) {
	reader := bufio.NewReader(c.CommandConn)
	line, _, err := reader.ReadLine()
	if err != nil {
		return
	}

	tokens := strings.Split(string(line), " ")
	command = tokens[0]
	args = tokens[1:]
	return
}

func (c *Connection) close() error {
	return c.CommandConn.Close()
}

func (c *Connection) reply(message string) error {
	_, err := c.CommandConn.Write([]byte(message))
	return err
}
