package charcount

import (
	"strings"
	"testing"
	"unicode/utf8"
)

func TestCountCharaters(t *testing.T) {
	var tests = []struct {
		input  string
		count  map[rune]int
		utflen [utf8.UTFMax + 1]int
	}{
		{"", make(map[rune]int), [utf8.UTFMax + 1]int{0, 0, 0, 0, 0}},
		{"a", map[rune]int{'a': 1}, [utf8.UTFMax + 1]int{0, 1, 0, 0, 0}},
		{"aa", map[rune]int{'a': 2}, [utf8.UTFMax + 1]int{0, 2, 0, 0, 0}},
		{"ab", map[rune]int{'a': 1, 'b': 1}, [utf8.UTFMax + 1]int{0, 2, 0, 0, 0}},
		{"すもももももももものうち", map[rune]int{'す': 1, 'も': 8, 'の': 1, 'う': 1, 'ち': 1}, [utf8.UTFMax + 1]int{0, 0, 0, 12, 0}},
		{"あああabab!", map[rune]int{'あ': 3, 'a': 2, 'b': 2, '!': 1}, [utf8.UTFMax + 1]int{0, 5, 0, 3, 0}},
	}
	for _, test := range tests {
		reader := strings.NewReader(test.input)
		count, utflen, err := CountCharacters(reader)
		if err != nil {
			t.Errorf("CountCharacters(%q) returned error: %v", test.input, err)
		} else {
			for key := range test.count {
				value, ok := count[key]
				if !ok {
					t.Errorf(
						"Result of CountCharacters(%q) should have key %c, but it actually don't have it",
						test.input, key)
				}
				if value != test.count[key] {
					t.Errorf(
						"Result of CountCharacters(%q) has key %c, but the count does't match. want %d, got %d",
						test.input, key, test.count[key], count[key])
				}
			}
			for key := range count {
				_, ok := test.count[key]
				if !ok {
					t.Errorf(
						"Result of CountCharacters(%q) contains unexpected key %c",
						test.input, key)
				}
			}

			if utflen != test.utflen {
				t.Errorf(
					"utflen of CountCharacters(%q) does not match with expected one. want %v, got %v",
					test.input, test.utflen, utflen)
			}
		}
	}
}
