package main_test

import (
	"testing"

	main "bootcamp/ch04/ex4.4"
)

func TestRotate(t *testing.T) {
	testcase := []struct {
		in  []int
		k   int
		out []int
	}{
		{[]int{1, 2, 3, 4, 5}, 2, []int{4, 5, 1, 2, 3}},
		{[]int{1, 2, 3, 4, 5}, 3, []int{3, 4, 5, 1, 2}},
		{[]int{1, 2, 3, 4, 5}, 4, []int{2, 3, 4, 5, 1}},
		{[]int{1, 2, 3, 4, 5}, 5, []int{1, 2, 3, 4, 5}},
		{[]int{1, 2, 3, 4, 5}, 6, []int{5, 1, 2, 3, 4}},
	}

	for _, c := range testcase {
		in := main.Rotate(c.in, c.k)

		for i, v := range in {
			if v != c.out[i] {
				t.Errorf("Rotate(%v, %d) = %v, want %v", c.in, c.k, in, c.out)
			}
		}
	}
}
