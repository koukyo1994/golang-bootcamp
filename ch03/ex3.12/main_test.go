package main_test

import (
	main "bootcamp/ch03/ex3.12"
	"testing"
)

func TestIsAnagram(t *testing.T) {
	testcase := []struct {
		s0   string
		s1   string
		want bool
	}{
		{"anagram", "naagram", true},
		{"anagram", "anagra", false},
		{"true", "eurt", true},
		{"boolian", "ian", false},
		{"", "", true},
		{"", "a", false},
		{"a", "", false},
	}

	for _, c := range testcase {
		if got := main.IsAnagram(c.s0, c.s1); got != c.want {
			t.Errorf("IsAnagram(%q, %q) = %v, want %v", c.s0, c.s1, got, c.want)
		}
	}
}
