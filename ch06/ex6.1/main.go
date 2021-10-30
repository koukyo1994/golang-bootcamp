package main

import (
	"bootcamp/ch06/ex6.1/intset"
	"fmt"
)

func main() {
	var x intset.IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)

	fmt.Printf("number of elements: %d\n", x.Len())
	fmt.Printf("x: %s\n", x.String())

	x.Remove(9)

	fmt.Printf("number of elements: %d\n", x.Len())
	fmt.Printf("x: %s\n", x.String())

	x.Clear()
	fmt.Printf("number of elements: %d\n", x.Len())
	fmt.Printf("x: %s\n", x.String())

	x.Add(42)
	x.Add(12)
	x.Add(144)

	y := x.Copy()

	y.Add(55)
	y.Add(25)

	fmt.Printf("x: %s\n", x.String())
	fmt.Printf("y: %s\n", y.String())
}
