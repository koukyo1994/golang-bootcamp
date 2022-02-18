package main

import "log"

func (c *Connection) handlels(args []string) {
	c.reply(StatusEnteringPassiveMode)

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
		c.reply(StatusDirectoryListing)
		c.reply(string(dirInfo))
	}
}

func (c *Connection) handlecd(args []string) {
	var to string
	if len(args) == 0 {
		to = string(c.RootDirectory)
	} else {
		to = args[0]
	}
	if err := c.WorkingDirectory.cd(to); err != nil {
		log.Print(err)
	} else {
		c.reply(StatusDirectoryChanged)
	}
}

func (c *Connection) handlepwd() {
	c.reply(StatusDisplayCurrentDirectory)
	c.reply(c.WorkingDirectory.pwd())
	c.reply("\n")
}

func (c *Connection) handlequit() {
	c.reply((StatusLogout))
	c.close()
}
