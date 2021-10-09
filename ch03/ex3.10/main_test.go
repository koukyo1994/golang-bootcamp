package main_test

import (
	"testing"

	main "bootcamp/ch03/ex3.10"
)

func TestComma(t *testing.T) {
	testcase := []struct {
		input string
		want  string
	}{
		{"12345", "12,345"},
		{"1234567", "1,234,567"},
		{"1234567890", "1,234,567,890"},
	}
	for _, c := range testcase {
		if got := main.Comma(c.input); got != c.want {
			t.Errorf("Comma(%q) = %q, want %q", c.input, got, c.want)
		}
	}
}
