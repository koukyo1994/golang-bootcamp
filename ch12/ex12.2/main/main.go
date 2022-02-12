package main

import "bootcamp/ch12/ex12.2/format"

type Cycle struct {
	Value int
	Tail  *Cycle
}

func main() {
	var c Cycle
	c = Cycle{42, &c}
	format.Display("c", c)
}
