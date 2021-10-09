package main

import (
	"bytes"
)

func Comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	b := []byte(s)
	var buf bytes.Buffer
	for i := 0; i < n; i++ {
		if i > 0 && (n-i)%3 == 0 {
			buf.WriteByte(',')
		}
		buf.WriteByte(b[i])
	}
	return buf.String()
}

func main() {
	s := "1234567"
	println(Comma(s))
}
