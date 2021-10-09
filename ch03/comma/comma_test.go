package comma_test

import "testing"

func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}

func TestComma(t *testing.T) {
	testcase := []struct {
		input string
		want  string
	}{
		{"123", "123"},
		{"1234", "1,234"},
		{"12345", "12,345"},
		{"123456", "123,456"},
		{"1234567", "1,234,567"},
		{"12345678", "12,345,678"},
		{"123456789", "123,456,789"},
		{"1234567890", "1,234,567,890"},
	}
	for _, c := range testcase {
		got := comma(c.input)
		if got != c.want {
			t.Errorf("comma(%q) == %q, want %q", c.input, got, c.want)
		}
	}
}
