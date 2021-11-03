package main

import "fmt"

func VariadicStringJoin(sep string, sentences ...string) string {
	var result string
	for i, sentence := range sentences {
		if i > 0 {
			result += sep
		}
		result += sentence
	}
	return result
}

func main() {
	fmt.Println(VariadicStringJoin(", ", "hello", "world"))
	fmt.Println(VariadicStringJoin("/", "path", "to", "somewhere"))
	fmt.Println(VariadicStringJoin(":", "onlyOneArgument"))
}
