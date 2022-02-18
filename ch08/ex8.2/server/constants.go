package main

const port = 5000

const (
	StatusEnteringPassiveMode     = "227 Entering Passive Mode\n"
	StatusDirectoryListing        = "150 Here comes the directory listing.\n"
	StatusDirectoryChanged        = "250 Directory successfully changed.\n"
	StatusDisplayCurrentDirectory = "257 "
)
