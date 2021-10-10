package main

import (
	"crypto/sha256"
	"fmt"
)

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func CountBitDifference(x, y [32]uint8) int {
	var count int
	for i := 0; i < len(x); i++ {
		xor := x[i] ^ y[i]
		count += int(pc[xor])
	}
	return count
}

func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	fmt.Printf("%x\n%x\n%d\n", c1, c2, CountBitDifference(c1, c2))
}
