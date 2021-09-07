package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	duplicatedFiles := make(map[string][]string)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, "Stdin", counts, duplicatedFiles)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, arg, counts, duplicatedFiles)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\t%s\n", n, line, strings.Join(duplicatedFiles[line], " "))
		}
	}
}

func countLines(f *os.File, fileName string, counts map[string]int, duplicatedFiles map[string][]string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
		if contains(duplicatedFiles[input.Text()], fileName) {
			continue
		}
		duplicatedFiles[input.Text()] = append(duplicatedFiles[input.Text()], fileName)
	}
	// 注意: input.Err()からのエラーの可能性を無視している
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
