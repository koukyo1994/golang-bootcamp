package eval

import (
	"fmt"
	"math"
	"strings"
)

// 配列に対する演算を表す。例: [0.2, 0.3, 0.4, 1.2].min
type arrayOp struct {
	fn    string // "min"、"max"、"sum"のどれか
	array []Expr
}

func (a arrayOp) Eval(env Env) float64 {
	switch a.fn {
	case "min":
		// +Inf
		min := math.Inf(1)
		for _, value := range a.array {
			min = math.Min(min, value.Eval(env))
		}
		return min
	case "max":
		// -Inf
		max := math.Inf(-1)
		for _, value := range a.array {
			max = math.Max(max, value.Eval(env))
		}
		return max
	case "sum":
		sum := 0.0
		for _, value := range a.array {
			sum += value.Eval(env)
		}
		return sum
	}
	panic(fmt.Sprintf("unsupported operation: %s", a.fn))
}

func isin(s string, set []string) bool {
	for _, elem := range set {
		if s == elem {
			return true
		}
	}
	return false
}

func (a arrayOp) Check(vars map[Var]bool) error {
	availableOp := []string{"min", "max", "sum"}
	if !isin(a.fn, availableOp) {
		return fmt.Errorf("unexpected array op %q", a.fn)
	}
	if len(a.array) == 0 {
		return fmt.Errorf("arrayOp requires more than one element in array, got %d", len(a.array))
	}
	for _, elem := range a.array {
		if err := elem.Check(vars); err != nil {
			return err
		}
	}
	return nil
}

func (a arrayOp) String() string {
	elements := []string{}
	for _, elem := range a.array {
		elements = append(elements, elem.String())
	}
	return fmt.Sprintf("[%s].%s()", strings.Join(elements, ", "), a.fn)
}
