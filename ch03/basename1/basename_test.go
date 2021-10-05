package basename1_test

import "testing"

func basename(s string) string {
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}
	return s
}

func TestBasename(testing *testing.T) {
	tests := []struct {
		path, want string
	}{
		{"a/b/c", "c"},
		{"a/b/c.go", "c"},
		{"a/b/c.h", "c"},
		{"a/", ""},
		{"a", "a"},
		{"/a/b/c", "c"},
		{"/", ""},
		{"", ""},
	}

	for _, test := range tests {
		if got := basename(test.path); got != test.want {
			testing.Errorf("basename(%q) = %q, want %q", test.path, got, test.want)
		}
	}
}
