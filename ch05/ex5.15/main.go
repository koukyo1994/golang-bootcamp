package main

import "fmt"

func min(vals ...int) int {
	min := vals[0]
	for _, v := range vals {
		if v < min {
			min = v
		}
	}
	return min
}

func max(vals ...int) int {
	max := vals[0]
	for _, v := range vals {
		if v > max {
			max = v
		}
	}
	return max
}

func minWithMoreThanOneArgs(val int, vals ...int) int {
	min := val
	for _, v := range vals {
		if v < min {
			min = v
		}
	}
	return min
}

func maxWithMoreThanOneArgs(val int, vals ...int) int {
	max := val
	for _, v := range vals {
		if v > max {
			max = v
		}
	}
	return max
}

func main() {
	fmt.Println(min(1, 2, 3, 4, 5))
	fmt.Println(max(1, 2, 3, 4, 5))
	// fmt.Println(min()) panic: index out of range
	// fmt.Println(max()) panic: index out of range
	fmt.Println(minWithMoreThanOneArgs(1, 2, 3, 4, 5))
	fmt.Println(maxWithMoreThanOneArgs(1, 2, 3, 4, 5))
	// fmt.Println(minWithMoreThanOneArgs()) compile error
	// fmt.Println(maxWithMoreThanOneArgs()) compile error
}
