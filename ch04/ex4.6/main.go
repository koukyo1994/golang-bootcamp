package main

import "unicode"

func NoAdjacentUnicodeSpace(slice []byte) []byte {
	i := 1
	for j, s := range slice {
		if j == 0 {
			continue
		}

		if unicode.IsSpace(rune(s)) {
			if unicode.IsSpace(rune(slice[j-1])) {
				continue
			}
			slice[i] = ' '
			i++
		} else {
			slice[i] = slice[j]
			i++
		}
	}
	return slice[:i]
}

func main() {
	s := []byte("a b c   d e f g  h i j k l   m n o p q r s t u v w x y z")
	s = NoAdjacentUnicodeSpace(s)
	println(string(s))
}
