package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

func main() {
	counts := make(map[rune]int)    // Unicode文字の数を数える
	var utflen [utf8.UTFMax + 1]int // UTF-8円コーディングの長さの数
	invalid := 0                    // 不正なUTF-8文字の数

	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune() // 1文字読み込む
		if err == io.EOF {
			break
		}
		if err != nil {
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		counts[r]++
		utflen[n]++
	}
	fmt.Printf("rune\tcount\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}
	fmt.Print("\nlen\tcount\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}
}
