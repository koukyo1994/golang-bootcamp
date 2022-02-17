package main

import "fmt"

type Function func(x int) int

func (f Function) After(g Function) Function {
	return func(x int) int { return g(f(x)) }
}

func f(x int) int { return x + 5 }
func g(x int) int { return x * x }

func main() {
	h := Function(f).After(g)
	fmt.Printf("%d\n", h(10))
}
