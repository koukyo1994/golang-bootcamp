package main

import "unicode/utf8"

func ReverseByteString(slice []byte) []byte {
	var head, tail int
	for {
		tail++
		// 文字ごとにreverseする
		if utf8.Valid(slice[head:tail]) {
			reverse(slice[head:tail])
			head = tail
		}
		if tail == len(slice) {
			break
		}
	}
	// 最後に全体をreverseする
	reverse(slice)
	return slice
}

func reverse(slice []byte) {
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
}

func main() {
	s := []byte("Hello, World!")
	s = ReverseByteString(s)
	println(string(s))

	s = []byte("Hello, 世界")
	s = ReverseByteString(s)
	println(string(s))

	s = []byte("這是一個 utf8 字符串的例子")
	s = ReverseByteString(s)
	println(string(s))
}
