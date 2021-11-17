package main

import (
	"regexp"
	"strconv"
	"strings"
	"testing"
)

func expand(s string, f func(string) string) string {
	reg := regexp.MustCompile(`\$\S*`)
	return reg.ReplaceAllStringFunc(s, func(s string) string {
		return f(s[1:])
	})
}

func TestExpand(t *testing.T) {
	tests := []struct {
		in       string
		f        func(string) string
		expected string
	}{
		{"aiu$eo kakiku", strings.ToUpper, "aiuEO kakiku"},
		{"abc$125 def$512", func(s string) string {
			value, _ := strconv.Atoi(s)
			return strconv.Itoa(value + 100)
		}, "abc225 def612"},
		{"", strings.Title, ""},
		{"aiueo", strings.ToTitle, "aiueo"},
	}
	for _, test := range tests {
		if got := expand(test.in, test.f); got != test.expected {
			t.Errorf("expand(%s, %p) = %s, want %s", test.in, &test.f, got, test.expected)
		}
	}
}
