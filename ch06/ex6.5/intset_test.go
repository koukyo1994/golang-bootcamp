package intset_test

import (
	intset "bootcamp/ch06/ex6.5"
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
