package cycle_test

import (
	"testing"

	cycle "bootcamp/ch13/ex13.2"
)

func TestHasCycle(t *testing.T) {
	type SampleStruct struct {
		value int
		ptr   *SampleStruct
	}
	a := SampleStruct{0.0, nil}
	b := SampleStruct{1.0, &a}
	var c SampleStruct
	c = SampleStruct{2.0, &c}

	cases := []struct {
		s    SampleStruct
		want bool
	}{
		{a, false},
		{b, false},
		{c, true},
	}

	for _, c := range cases {
		if cycle.HasCycle(c.s) != c.want {
			t.Errorf("HasCycle(%v) = %v, want %v", c.s, !c.want, c.want)
		}
	}
}
