package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()

	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)

	secs := time.Since(start).Seconds()
	msg := fmt.Sprintf("非効率な書き方: %.6fs", secs)
	fmt.Println(msg)

	start = time.Now()
	fmt.Println(strings.Join(os.Args[1:], " "))
	secs = time.Since(start).Seconds()
	msg = fmt.Sprintf("strings.Joinを使った書き方: %6fs", secs)
	fmt.Println(msg)
}
