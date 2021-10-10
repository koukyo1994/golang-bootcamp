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
	unicodeClassification := map[string]int{
		"Control": 0,
		"Digit":   0,
		"Letter":  0,
		"Graphic": 0,
		"Punct":   0,
		"Space":   0,
		"Mark":    0,
	}
	invalid := 0 // 不正なUTF-8文字の数

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

		if unicode.IsControl(r) {
			unicodeClassification["Control"]++
		}
		if unicode.IsDigit(r) {
			unicodeClassification["Digit"]++
		}
		if unicode.IsLetter(r) {
			unicodeClassification["Letter"]++
		}
		if unicode.IsGraphic(r) {
			unicodeClassification["Graphic"]++
		}
		if unicode.IsPunct(r) {
			unicodeClassification["Punct"]++
		}
		if unicode.IsSpace(r) {
			unicodeClassification["Space"]++
		}
		if unicode.IsMark(r) {
			unicodeClassification["Mark"]++
		}
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
	fmt.Print("\nUnicode\tClassification\n")
	for c, n := range unicodeClassification {
		fmt.Printf("%s\t%d\n", c, n)
	}
}
