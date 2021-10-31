package intset_test

import (
	intset "bootcamp/ch06/ex6.3"
	"testing"
)

func TestIntersectWith(t *testing.T) {
	cases := []struct {
		swords []int
		twords []int
		want   string
	}{
		{[]int{1, 2, 3}, []int{1, 2, 3}, "{1 2 3}"},
		{nil, []int{1, 2, 3}, "{}"},
		{[]int{1, 2, 3}, nil, "{}"},
		{[]int{1, 2, 3}, []int{4, 5, 6}, "{}"},
		{[]int{1, 2, 3}, []int{1, 2, 3, 4, 5, 6}, "{1 2 3}"},
		{[]int{1, 2, 3}, []int{100, 142, 200}, "{}"},
		{[]int{1, 2, 3, 142, 255, 368}, []int{1, 142, 12}, "{1 142}"},
	}
	for _, c := range cases {
		source := intset.IntSet{}
		target := intset.IntSet{}

		source.AddAll(c.swords...)
		target.AddAll(c.twords...)

		if source.IntersectWith(&target); source.String() != c.want {
			t.Errorf("IntersectWith(%v, %v) = %v, want %v", c.swords, c.twords, source.String(), c.want)
		}
	}
}

func TestDifferenceWith(t *testing.T) {
	cases := []struct {
		swords []int
		twords []int
		want   string
	}{
		{[]int{1, 2, 3}, []int{1, 2, 3}, "{}"},
		{nil, []int{1, 2, 3}, "{}"},
		{[]int{1, 2, 3}, nil, "{1 2 3}"},
		{[]int{1, 2, 3}, []int{4, 5, 6}, "{1 2 3}"},
		{[]int{1, 2, 3}, []int{1, 2, 3, 4, 5, 6}, "{}"},
		{[]int{1, 2, 3, 144, 999, 568}, []int{1, 144, 12}, "{2 3 568 999}"},
	}
	for _, c := range cases {
		source := intset.IntSet{}
		target := intset.IntSet{}

		source.AddAll(c.swords...)
		target.AddAll(c.twords...)

		if source.DifferenceWith(&target); source.String() != c.want {
			t.Errorf("DifferenceWith(%v, %v) = %v, want %v", c.swords, c.twords, source.String(), c.want)
		}
	}
}

func TestSymmetricDifference(t *testing.T) {
	cases := []struct {
		swords []int
		twords []int
		want   string
	}{
		{[]int{1, 2, 3}, []int{1, 2, 3}, "{}"},
		{nil, []int{1, 2, 3}, "{1 2 3}"},
		{[]int{1, 2, 3}, nil, "{1 2 3}"},
		{[]int{1, 2, 3}, []int{4, 5, 6}, "{1 2 3 4 5 6}"},
		{[]int{1, 2, 3}, []int{1, 2, 3, 4, 5, 6}, "{4 5 6}"},
		{[]int{1, 2, 3, 144, 999, 568}, []int{1, 144, 12}, "{2 3 12 568 999}"},
	}
	for _, c := range cases {
		source := intset.IntSet{}
		target := intset.IntSet{}

		source.AddAll(c.swords...)
		target.AddAll(c.twords...)

		if source.SymmetricDifference(&target); source.String() != c.want {
			t.Errorf("SymmetricDifference(%v, %v) = %v, want %v", c.swords, c.twords, source.String(), c.want)
		}
	}
}
