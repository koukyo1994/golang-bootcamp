package main_test

import (
	popcount "bootcamp/ch02/popcount"
	"testing"
)

var count int

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		count = popcount.PopCount(uint64(i))
	}
}

func BenchmarkPopCountLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		count = popcount.PopCountLoop(uint64(i))
	}
}

func BenchmarkPopCountShift(b *testing.B) {
	for i := 0; i < b.N; i++ {
		count = popcount.PopCountShift(uint64(i))
	}
}

func BenchmarkPopCountClear(b *testing.B) {
	for i := 0; i < b.N; i++ {
		count = popcount.PopCountClear(uint64(i))
	}
}
