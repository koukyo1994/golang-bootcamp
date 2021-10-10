package main_test

import (
	"testing"

	main "bootcamp/ch04/ex4.6"
)

func TestNoAdjacentUnicodeSpace(t *testing.T) {
	testcase := []struct {
		in       []byte
		expected []byte
	}{
		{
			[]byte("a b c"),
			[]byte("a b c"),
		},
		{
			[]byte("a b c   d e f g  h i j k l   m n o p q r s t u v w x y z"),
			[]byte("a b c d e f g h i j k l m n o p q r s t u v w x y z"),
		},
	}

	for _, c := range testcase {
		actual := main.NoAdjacentUnicodeSpace(c.in)
		if string(actual) != string(c.expected) {
			t.Errorf("NoAdjacentUnicodeSpace(%q) == %q, expected %q", c.in, actual, c.expected)
		}
	}
}
