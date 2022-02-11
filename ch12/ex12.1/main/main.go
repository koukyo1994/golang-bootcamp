package main

import "bootcamp/ch12/ex12.1/format"

type Key struct {
	a, b int
	c    string
}

type Sample struct {
	d map[Key]string
	e map[[2]int]int
}

func main() {
	sample := Sample{
		d: map[Key]string{
			{a: 1, b: 2, c: "aaa"}: "aaa",
			{a: 2, b: 4, c: "bbb"}: "bbb",
		},
		e: map[[2]int]int{
			{1, 2}: 3,
			{2, 3}: 4,
		},
	}
	format.Display("sample", sample)
}
