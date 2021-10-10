package rev_test

import (
	"testing"

	main "bootcamp/ch04/ex4.3"
)

func TestReverse(t *testing.T) {
	testcase := []struct {
		input []int
		want  []int
	}{
		{[]int{1, 2, 3}, []int{3, 2, 1}},
		{[]int{1, 2, 3, 4}, []int{4, 3, 2, 1}},
		{[]int{1, 2, 3, 4, 5}, []int{5, 4, 3, 2, 1}},
	}

	for _, c := range testcase {
		in := c.input
		main.Reverse(&in)

		match := true
		for i, v := range in {
			if v != c.want[i] {
				match = false
				break
			}
		}
		if !match {
			t.Errorf("Reverse(%v) == %v, want %v", c.input, in, c.want)
		}
	}
}
