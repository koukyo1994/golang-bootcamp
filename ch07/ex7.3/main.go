package main

import (
	"fmt"
	"strconv"
)

type tree struct {
	value       int
	left, right *tree
}

// Sortはvalues内の値をその中でソートします
func Sort(values []int) *tree {
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}
	appendValues(values[:0], root)
	return root
}

// appendValuesはtの要素をvaluesの正しい順序に追加し、結果のスライスを返します
func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

func add(t *tree, value int) *tree {
	if t == nil {
		// return &tree{value: value}と同じ
		t = new(tree)
		t.value = value
		return t
	}
	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}

func (t *tree) String() string {
	if t.left == nil && t.right == nil {
		return strconv.Itoa(t.value)
	}
	s := strconv.Itoa(t.value)
	if t.right != nil {
		s += "--r--" + t.right.String() + "\n"
	} else {
		s += "\n"
	}
	if t.left != nil {
		s += "|\n"
		s += "l\n"
		s += "|\n"
		s += t.left.String() + "\n"
	}
	return s
}

func main() {
	values := []int{5, 3, 6, 2, 4, 1}
	root := Sort(values)
	fmt.Printf("%s", root)

	values = []int{10, 39, 22, 48, 67, 17, 2, 9, 33, 6, 52, 34, 33, 91, 1, 24}
	root = Sort(values)
	fmt.Printf("%s", root)
}
