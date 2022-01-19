package charcount

import (
	"bufio"
	"io"
	"unicode"
	"unicode/utf8"
)

func CountCharacters(reader io.Reader) (counts map[rune]int, utflen [utf8.UTFMax + 1]int, err error) {
	in := bufio.NewReader(reader)
	counts = make(map[rune]int)
	for {
		r, n, err := in.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			return counts, utflen, err
		}
		if r == unicode.ReplacementChar && n == 1 {
			continue
		}
		counts[r]++
		utflen[n]++
	}
	return
}
