package main

import "fmt"

func NoAdjacent(strings []string) []string {
	i := 1
	for j, s := range strings {
		if j == 0 {
			continue
		}

		if s != strings[j-1] {
			strings[i] = s
			i++
		}
	}
	return strings[:i]
}

func main() {
	strings := []string{"a", "b", "a", "a", "a", "c", "c", "d", "d", "d"}
	strings = NoAdjacent(strings)
	fmt.Printf("%q\n", strings)

	strings = []string{"a", "b", "c", "d", "e", "e", "f", "f", "f", "g"}
	strings = NoAdjacent(strings)
	fmt.Printf("%q\n", strings)
}
