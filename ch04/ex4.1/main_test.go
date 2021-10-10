package main_test

import (
	"testing"

	main "bootcamp/ch04/ex4.1"
)

func TestCountBitDifference(t *testing.T) {
	testcase := []struct {
		input0   [32]uint8
		input1   [32]uint8
		expected int
	}{
		{
			[32]uint8{},
			[32]uint8{},
			0,
		}, {
			[32]uint8{31: 1},
			[32]uint8{30: 1, 31: 1},
			1,
		}, {
			[32]uint8{},
			[32]uint8{255},
			8,
		},
	}

	for _, c := range testcase {
		actual := main.CountBitDifference(c.input0, c.input1)
		if actual != c.expected {
			t.Errorf("CountBitDifference(%v, %v) = %v, expected %v", c.input0, c.input1, actual, c.expected)
		}
	}
}
