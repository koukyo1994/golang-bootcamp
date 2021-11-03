package main_test

import (
	main "bootcamp/ch05/ex5.16"
	"testing"
)

func TestVariadicStringJoin(t *testing.T) {
	var tests = []struct {
		sep  string
		str  []string
		want string
	}{
		{"", []string{"a"}, "a"},
		{"", []string{"a", "b"}, "ab"},
		{",", []string{"a", "b", "c"}, "a,b,c"},
		{"/", []string{"a", "b", "c", "d"}, "a/b/c/d"},
	}
	for _, test := range tests {
		if got := main.VariadicStringJoin(test.sep, test.str...); got != test.want {
			t.Errorf("VariadicStringJoin(%q, %q) = %q, want %q", test.sep, test.str, got, test.want)
		}
	}
}
