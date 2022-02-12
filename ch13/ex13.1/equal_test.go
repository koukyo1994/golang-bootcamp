package equal_test

import (
	"math"
	"testing"

	eq "bootcamp/ch13/ex13.1"
)

func TestEqual(t *testing.T) {
	intcases := []struct {
		input0 int
		input1 int
		want   bool
	}{
		{0, 0, true},
		{0, 1, false},
		{1e8, 1e8 + 1, false},
		{1e9, 1e9, true},
		{-1e9, 1e9, false},
		{math.MaxInt64, math.MaxInt64, true},
		{math.MaxInt64, -math.MaxInt64, false},
	}
	for _, c := range intcases {
		if eq.Equal(c.input0, c.input1) != c.want {
			t.Errorf("Equal(%d, %d) = %v, want %v", c.input0, c.input1, !c.want, c.want)
		}
	}

	uintcases := []struct {
		input0 uint
		input1 uint
		want   bool
	}{
		{0, 0, true},
		{0, 1, false},
		{math.MaxUint64, math.MaxUint64, true},
		{math.MaxUint64, 0, false},
	}
	for _, c := range uintcases {
		if eq.Equal(c.input0, c.input1) != c.want {
			t.Errorf("Equal(%d, %d) = %v, want %v", c.input0, c.input1, !c.want, c.want)
		}
	}

	floatcases := []struct {
		input0 float64
		input1 float64
		want   bool
	}{
		{0.0, 0.0, true},
		{0.0, 1.0, false},
		{math.MaxFloat64, math.MaxFloat64, true},
		{math.MaxFloat64, -math.MaxFloat64, false},
		{0.0 - 9e-10, 0.0, true},
		{0.0 - 1e-9, 0.0, false},
	}
	for _, c := range floatcases {
		if eq.Equal(c.input0, c.input1) != c.want {
			t.Errorf("Equal(%f, %f) = %v, want %v", c.input0, c.input1, !c.want, c.want)
		}
	}

	cmplxcases := []struct {
		input0 complex128
		input1 complex128
		want   bool
	}{
		{0.0 + 1i, 0.0 + 1i, true},
		{1.0 + 1i, 0.0 + 1i, false},
		{0.0 + 1i, 0.0 + 2i, false},
		{5e-10 + 1i, 0.0 + 1i, true},
		{1e-9 + 1i, 0.0 + 1i, false},
		{5e-10 + (5e-10+1)*1i, 0.0 + 1i, true},
		{1e-10 + (5e-10+1)*1i, 0.0 + 1i, true},
		{9e-10 + (5e-10+1)*1i, 0.0 + 1i, false},
	}
	for _, c := range cmplxcases {
		if eq.Equal(c.input0, c.input1) != c.want {
			t.Errorf("Equal(%v, %v) = %v, want %v", c.input0, c.input1, !c.want, c.want)
		}
	}
}
