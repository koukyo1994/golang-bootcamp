package intset_test

import (
	intset "bootcamp/ch06/ex6.2"
	"testing"
)

func TestAddAll(t *testing.T) {
	cases := []struct {
		words []int
		add   []int
		want  string
	}{
		{[]int{1, 2, 3}, []int{4, 5, 6}, "{1 2 3 4 5 6}"},
		{nil, []int{1, 2, 3}, "{1 2 3}"},
		{[]int{2, 5, 8}, []int{1, 3, 4, 6, 7, 9}, "{1 2 3 4 5 6 7 8 9}"},
		{[]int{1, 2, 3}, []int{1}, "{1 2 3}"},
		{[]int{1, 2, 3}, nil, "{1 2 3}"},
	}
	for _, c := range cases {
		s := intset.IntSet{}
		for _, w := range c.words {
			s.Add(w)
		}
		s.AddAll(c.add...)
		if s.String() != c.want {
			t.Errorf("AddAll(%v) = %s, want %s", c.add, s.String(), c.want)
		}
	}
}
