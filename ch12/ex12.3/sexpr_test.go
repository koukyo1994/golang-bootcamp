package sexpr_test

import (
	"testing"

	sexpr "bootcamp/ch12/ex12.3"
)

func TestMarshal(t *testing.T) {
	boolcases := []struct {
		input bool
		want  string
	}{
		{true, "t"},
		{false, "nil"},
	}

	for _, c := range boolcases {
		got, err := sexpr.Marshal(c.input)
		if err != nil {
			t.Errorf("Marshal(%v) returned error: %v", c.input, err)
		} else {
			if string(got) != c.want {
				t.Errorf("Marshal(%v) = %s, want %s", c.input, string(got), c.want)
			}
		}
	}

	floatcases := []struct {
		input float64
		want  string
	}{
		{3.141592, "3.141592"},
		{0.0, "0"},
		{-300.0331298, "-300.0331298"},
	}

	for _, c := range floatcases {
		got, err := sexpr.Marshal(c.input)
		if err != nil {
			t.Errorf("Marshal(%v) returned error: %v", c.input, err)
		} else {
			if string(got) != c.want {
				t.Errorf("Marshal(%v) = %s, want %s", c.input, string(got), c.want)
			}
		}
	}

	cmplxcases := []struct {
		input complex128
		want  string
	}{
		{1.0 + 3.0i, "#C(1 3)"},
		{0.0 + 1.9i, "#C(0 1.9)"},
		{-3.9 - 2.1i, "#C(-3.9 -2.1)"},
	}

	for _, c := range cmplxcases {
		got, err := sexpr.Marshal(c.input)
		if err != nil {
			t.Errorf("Marshal(%v) returned error: %v", c.input, err)
		} else {
			if string(got) != c.want {
				t.Errorf("Marshal(%v) = %s, want %s", c.input, string(got), c.want)
			}
		}
	}
}
