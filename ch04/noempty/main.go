package main

import "fmt"

// noemptyは空文字列でない文字列を保持するスライスを返す。
// 基底配列は呼び出し中に修正される。
func noempty(strings []string) []string {
	i := 0
	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i++
		}
	}
	return strings[:i]
}

func noempty2(strings []string) []string {
	out := strings[:0] // 元の長さ0のスライス
	for _, s := range strings {
		if s != "" {
			out = append(out, s)
		}
	}
	return out
}

func main() {
	data := []string{"one", "", "three"}
	fmt.Printf("%q\n", noempty(data)) // "[one three]"
	fmt.Printf("%q\n", data)          // "[one three three]"

	data = []string{"one", "", "three"}
	fmt.Printf("%q\n", noempty2(data)) // "[one three]"
	fmt.Printf("%q\n", data)           // "[one three three]"
}
