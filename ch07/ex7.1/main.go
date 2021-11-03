package main

import (
	"bufio"
	"bytes"
)

type WordCounter int
type LineCounter int

func (w *WordCounter) Write(p []byte) (int, error) {
	reader := bytes.NewReader(p)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		*w++
	}
	return len(p), nil
}

func (l *LineCounter) Write(p []byte) (int, error) {
	reader := bytes.NewReader(p)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		*l++
	}
	return len(p), nil
}

func main() {
	var wc WordCounter
	var lc LineCounter
	wc.Write([]byte("hello world! This is sample sentences.\n Can you count how many words are in this byte buffer?\n"))
	lc.Write([]byte("hello world! This is sample sentences.\n Can you count how many lines are in this byte buffer?\n"))
	println(wc)
	println(lc)
}
