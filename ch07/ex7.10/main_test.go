package main

import (
	"sort"
	"testing"
)

func IsPalindrome(s sort.Interface) bool {
	result := true
	for i := 0; i < s.Len(); i++ {
		result = result && !s.Less(i, s.Len()-i-1) && !s.Less(s.Len()-i-1, i)
	}
	return result
}

type intSlice []int

func (x intSlice) Len() int           { return len(x) }
func (x intSlice) Less(i, j int) bool { return x[i] < x[j] }
func (x intSlice) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

type byteSlice []byte

func (x byteSlice) Len() int           { return len(x) }
func (x byteSlice) Less(i, j int) bool { return x[i] < x[j] }
func (x byteSlice) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

func TestIsPalindrome(t *testing.T) {
	cases := []struct {
		s        sort.Interface
		expected bool
	}{
		{intSlice([]int{1, 2, 3, 4, 5, 4, 3, 2, 1}), true},
		{intSlice([]int{1, 2, 3, 4, 5, 4, 3, 2, 2}), false},
		{byteSlice([]byte("aeiouoiea")), true},
		{byteSlice([]byte("aeiouuuiea")), false},
	}

	for _, c := range cases {
		if IsPalindrome(c.s) != c.expected {
			t.Errorf("IsPalindrome(%v) == %v, want %v", c.s, !c.expected, c.expected)
		}
	}
}
