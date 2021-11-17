package eval_test

import (
	eval "bootcamp/ch07/ex7.14"
	"fmt"
	"testing"
)

func TestArrayOP(t *testing.T) {
	tests := []struct {
		expr string
		env  eval.Env
		want string
	}{
		{"[x, y, 0.3].min", eval.Env{"x": 0.5, "y": 0.2, "z": 0.1}, "0.2"},
		{"[0.4, 0.8, A, 0.7].max", eval.Env{"A": 1.5}, "1.5"},
		{"[0.1, 1.2, 3.0, c].sum", eval.Env{"a": 0.1, "b": 4.0, "c": 2.0}, "6.3"},
	}
	for _, test := range tests {
		expr, err := eval.Parse(test.expr)
		if err != nil {
			t.Error(err) // パースエラー
			continue
		}
		got := fmt.Sprintf("%.6g", expr.Eval(test.env))
		fmt.Printf("\t%s in %v => %s\n", test.expr, test.env, got)
		if got != test.want {
			t.Errorf("%s.Eval() in %v = %q, want %q\n", test.expr, test.env, got, test.want)
		}
	}
}
