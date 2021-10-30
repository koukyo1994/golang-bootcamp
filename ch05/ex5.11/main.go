package main

import (
	"fmt"
)

var prereqs = map[string][]string{
	"algorithms":     {"data structures"},
	"calculus":       {"linear algebra"},
	"linear algebra": {"calculus"},
	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},
	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

const (
	temporal = 0
	settled  = 1
)

func topoSort(m map[string][]string) (sorted []string, ok bool) {
	var order []string
	statusMap := make(map[string]int)
	var visitAll func(key string) bool

	visitAll = func(item string) bool {
		status, ok := statusMap[item]
		if !ok {
			statusMap[item] = temporal
			for _, item := range m[item] {
				if ok = visitAll(item); !ok {
					return ok
				}
			}
			statusMap[item] = settled
			order = append(order, item)
		} else if status == temporal {
			return false
		}
		return true
	}

	for item := range m {
		if _, ok := statusMap[item]; !ok {
			ok = visitAll(item)
			if !ok {
				return nil, ok
			}
		}
	}
	return order, true
}

func main() {
	sorted, ok := topoSort(prereqs)
	if !ok {
		fmt.Printf("Cycle detected.\n")
	} else {
		for i, course := range sorted {
			fmt.Printf("%d:\t%s\n", i+1, course)
		}
	}
}
