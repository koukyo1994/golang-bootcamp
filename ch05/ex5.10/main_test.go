package main

import "testing"

func getIndex(item string, array []string) int {
	for i, elem := range array {
		if item == elem {
			return i
		}
	}
	return -1
}

func TestTopoSort(t *testing.T) {
	sorted := topoSort(prereqs)
	for key, values := range prereqs {
		indexOfKey := getIndex(key, sorted)
		if indexOfKey == -1 {
			t.Fatalf("results of `topoSort` seems to have deleted some elements of the input: %s not found", key)
		}
		for _, value := range values {
			indexOfValue := getIndex(value, sorted)
			if indexOfValue == -1 {
				t.Fatalf("results of `topoSort` seems to have deleted some elements of the input: %s not found", value)
			}
			if indexOfKey < indexOfValue {
				t.Fatalf("%s must be after %s, but the result is in contradictory order", key, value)
			}
		}
	}
}
