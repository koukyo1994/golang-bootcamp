package main

import (
	"strings"
)

func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}

func reverseString(s string) string {
	b := []byte(s)
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	return string(b)
}

func CommaForFloat(s string) string {
	splitted := strings.Split(s, ".")
	if len(splitted) == 1 {
		return comma(splitted[0])
	} else if len(splitted) == 2 {
		return comma(splitted[0]) + "." + reverseString(comma(reverseString(splitted[1])))
	} else {
		panic("invalid float")
	}
}

func main() {
	println(CommaForFloat("123456789"))
	println(CommaForFloat("123456789.123456789"))
	println(CommaForFloat("123456789.123456789123456789"))
}
