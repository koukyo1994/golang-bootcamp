package eval_test

import (
	eval "bootcamp/ch07/ex7.13"
	"testing"
)

func TestString(t *testing.T) {
	tests := []string{
		"sin(-x) * pow(1.5,-r)",
		"pow(2, sin(y)) * pow(2,sin(x)) / 12",
		"sin(x * y / 10)/10",
	}
	for _, test := range tests {
		expr, err := eval.Parse(test)
		if err != nil {
			t.Error(err)
			continue
		}

		stringified := expr.String()
		reexpr, err := eval.Parse(stringified)
		if err != nil {
			t.Fatalf("parsing %s: %v", stringified, err)
		}

		env := eval.Env{"x": 0.1, "y": 0.1, "r": 0.3}
		if expr.Eval(env) != reexpr.Eval(env) {
			t.Fatalf("%s.Eval(env) != %s.Eval(env)", test, stringified)
		}
	}
}
