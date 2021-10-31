package intset_test

import (
	intset "bootcamp/ch06/ex6.4"
	"testing"
)

func TestElems(t *testing.T) {
	cases := []struct {
		words []int
		want  []int
	}{
		{[]int{1, 2, 3}, []int{1, 2, 3}},
		{nil, nil},
		{[]int{5, 8, 2, 3, 1, 2, 10, 100, 200}, []int{1, 2, 3, 5, 8, 10, 100, 200}},
	}
	for _, c := range cases {
		s := intset.IntSet{}

		got := s.Elems()
		for i, v := range got {
			if c.want[i] != v {
				t.Errorf("got %v, want %v", got, c.want)
			}
		}
	}
}
