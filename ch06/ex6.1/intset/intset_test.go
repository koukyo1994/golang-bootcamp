package intset_test

import (
	"bootcamp/ch06/ex6.1/intset"
	"testing"
)

func TestLen(t *testing.T) {
	cases := []struct {
		words []int
		want  int
	}{
		{nil, 0},
		{[]int{1, 2, 3}, 3},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 10},
	}
	for _, c := range cases {
		s := intset.IntSet{}
		for _, w := range c.words {
			s.Add(w)
		}
		if got := s.Len(); got != c.want {
			t.Errorf("Len(%v) = %d, want %d", c.words, got, c.want)
		}
	}
}

func TestRemove(t *testing.T) {
	cases := []struct {
		words  []int
		remove int
		want   string
	}{
		{[]int{1, 2, 3}, 1, "{2 3}"},
		{[]int{1, 2, 3}, 2, "{1 3}"},
		{[]int{1}, 1, "{}"},
	}

	for _, c := range cases {
		s := intset.IntSet{}
		for _, w := range c.words {
			s.Add(w)
		}
		s.Remove(c.remove)
		if got := s.String(); got != c.want {
			t.Errorf("Remove(%d) = %s, want %s", c.remove, got, c.want)
		}
	}
}

func TestClear(t *testing.T) {
	cases := []struct{ words []int }{
		{[]int{1, 2, 3}},
		{nil},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
	}
	for _, c := range cases {
		s := intset.IntSet{}
		for _, w := range c.words {
			s.Add(w)
		}
		s.Clear()
		if got := s.Len(); got != 0 {
			t.Errorf("Clear() = %d, want 0", got)
		}
	}
}

func TestCopy(t *testing.T) {
	cases := []struct {
		words    []int
		add      int
		original string
		copy     string
	}{
		{[]int{1, 2, 3}, 1, "{1 2 3}", "{1 2 3}"},
		{[]int{1, 2, 3}, 4, "{1 2 3}", "{1 2 3 4}"},
		{nil, 1, "{}", "{1}"},
	}
	for _, c := range cases {
		s := intset.IntSet{}
		for _, w := range c.words {
			s.Add(w)
		}

		cpy := s.Copy()
		cpy.Add(c.add)
		if got := s.String(); got != c.original {
			t.Errorf("Original: %s, want %s", got, c.original)
		}
		if got := cpy.String(); got != c.copy {
			t.Errorf("Copy: %s, want %s", got, c.copy)
		}
	}
}
