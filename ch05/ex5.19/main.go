package main

import "fmt"

func f() (s string) {
	defer func() {
		s = recover().(string)
	}()
	panic("hello, world!")
}

func main() {
	s := f()
	fmt.Println(s)
}
