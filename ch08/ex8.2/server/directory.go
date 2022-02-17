package main

import (
	"os"
	"os/exec"
)

type Directory string

func (d *Directory) cd(to string) error {
	if err := os.Chdir(string(*d)); err != nil {
		return err
	}

	if err := os.Chdir(to); err != nil {
		return err
	}

	currentDir, err := os.Getwd()
	if err != nil {
		return err
	}

	*d = Directory(currentDir)
	return nil
}

func (d Directory) pwd() string {
	return string(d)
}

func (d Directory) ls(path string) ([]byte, error) {
	if err := os.Chdir(string(d)); err != nil {
		return nil, err
	}
	return exec.Command("ls", "-l", path).Output()
}
