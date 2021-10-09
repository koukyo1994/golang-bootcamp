package main

import "fmt"

func contains(s []rune, c rune) (bool, int) {
	for i, r := range s {
		if r == c {
			return true, i
		}
	}
	return false, -1
}

func IsAnagram(s0 string, s1 string) bool {
	if len(s0) != len(s1) {
		return false
	}
	r1 := []rune(s1)
	for _, r := range s0 {
		found, j := contains(r1, r)
		if !found {
			return false
		} else {
			r1 = append(r1[:j], r1[j+1:]...)
		}
	}
	return true
}

func main() {
	s0 := "anagram"
	s1 := "nagaram"
	if IsAnagram(s0, s1) {
		fmt.Printf("%s and %s are anagrams\n", s0, s1)
	} else {
		fmt.Printf("%s and %s are not anagrams\n", s0, s1)
	}
}
