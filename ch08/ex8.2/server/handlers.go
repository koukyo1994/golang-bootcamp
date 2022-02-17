package main

import "log"

func (c *Connection) handlels(args []string) {
	var path string
	if len(args) == 0 {
		path = "."
	} else {
		path = args[0]
	}
	dirInfo, err := c.WorkingDirectory.ls(path)
	if err != nil {
		log.Print(err)
	} else {
		log.Print(string(dirInfo))
	}
}
