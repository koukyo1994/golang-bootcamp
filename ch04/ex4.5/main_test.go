package main_test

import (
	"testing"

	main "bootcamp/ch04/ex4.5"
)

func TestNoAdjacent(t *testing.T) {
	testcase := []struct {
		in       []string
		expected []string
	}{
		{
			[]string{"a", "b", "a", "a", "a", "c", "c", "d", "d", "d"},
			[]string{"a", "b", "a", "c", "d"},
		},
		{
			[]string{"a", "b", "c", "d", "e", "e", "f", "f", "f", "g"},
			[]string{"a", "b", "c", "d", "e", "f", "g"},
		},
	}

	for _, c := range testcase {
		actual := main.NoAdjacent(c.in)

		if len(actual) != len(c.expected) {
			t.Errorf("NoAdjacent(%v) == %v, expected %v", c.in, actual, c.expected)
		}

		for i, v := range actual {
			if v != c.expected[i] {
				t.Errorf("NoAdjacent(%v) == %v, expected %v", c.in, actual, c.expected)
			}
		}
	}
}
