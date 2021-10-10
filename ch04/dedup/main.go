package main

import (
	"bufio"
	"os"
)

func main() {
	seen := make(map[string]bool)
	input := bufio.NewScanner(os.Stdin)

	for input.Scan() {
		line := input.Text()
		if !seen[line] {
			seen[line] = true
			println(line)
		}
	}

	if err := input.Err(); err != nil {
		println("error:", err)
		os.Exit(1)
	}
}
