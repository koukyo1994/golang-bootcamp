package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
)

type Dependencies struct {
	Deps []string
}

func getPackageDependencies(packages []string) ([]Dependencies, error) {
	allDeps := []Dependencies{}
	for _, packageName := range packages {
		command := exec.Command("go", "list", "-json", "-e", packageName)
		stdout := bytes.Buffer{}
		command.Stdout = &stdout
		err := command.Run()
		if err != nil {
			return nil, err
		}
		deps := Dependencies{}
		if err := json.NewDecoder(&stdout).Decode(&deps); err != nil {
			return nil, err
		}
		allDeps = append(allDeps, deps)
	}
	return allDeps, nil
}

func main() {
	packages, err := getPackageDependencies(os.Args[1:])
	if err != nil {
		fmt.Fprintf(os.Stderr, "searching dependencies: %v\n", err)
		os.Exit(1)
	}
	for _, p := range packages {
		for _, d := range p.Deps {
			fmt.Println(d)
		}
	}
}
