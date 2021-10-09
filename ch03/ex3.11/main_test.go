package main_test

import (
	"testing"

	main "bootcamp/ch03/ex3.11"
)

func TestCommaForFloat(t *testing.T) {
	testcase := []struct {
		input string
		want  string
	}{
		{"12345.67", "12,345.67"},
		{"12345.67899", "12,345.678,99"},
		{"123456789.987654321", "123,456,789.987,654,321"},
	}

	for _, c := range testcase {
		got := main.CommaForFloat(c.input)
		if got != c.want {
			t.Errorf("Comma(%q) == %q, want %q", c.input, got, c.want)
		}
	}
}
