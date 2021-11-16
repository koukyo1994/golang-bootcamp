package eval

import (
	"fmt"
	"strings"
)

func (v Var) String() string {
	return string(v)
}

func (l literal) String() string {
	return fmt.Sprint(float64(l))
}

func (u unary) String() string {
	if _, ok := u.x.(binary); ok {
		return fmt.Sprintf("%c(%s)", u.op, u.x)
	} else {
		return fmt.Sprintf("%c%s", u.op, u.x)
	}
}

func (b binary) String() string {
	var x, y string
	if _, ok := b.x.(binary); ok {
		x = fmt.Sprintf("(%s)", b.x)
	} else {
		x = b.x.String()
	}
	if _, ok := b.y.(binary); ok {
		y = fmt.Sprintf("(%s)", b.y)
	} else {
		y = b.y.String()
	}
	return fmt.Sprintf("%s %c %s", x, b.op, y)
}

func (c call) String() string {
	args := []string{}
	for _, arg := range c.args {
		args = append(args, arg.String())
	}
	return fmt.Sprintf("%s(%s)", c.fn, strings.Join(args, ", "))
}
