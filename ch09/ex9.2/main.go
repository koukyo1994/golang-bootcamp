package main

import (
	"fmt"
	"sync"
	"time"
)

var pcInit sync.Once

// pc[i]はiのポピュレーションカウント
var pc [256]byte

func calculateTable() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCountはxのポピュレーションカウント(1が設定されているビット数)を返します
func PopCount(x uint64) int {
	pcInit.Do(calculateTable)
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

func main() {
	var x uint64 = 0x1234567890ABCDEF
	fmt.Printf("%x\n", x)
	t0 := time.Now()
	fmt.Println(PopCount(x))
	elapsed := time.Since(t0)
	fmt.Printf("%.8f s elapsed\n", elapsed.Seconds())

	x = 0x234567890ABCDEF
	fmt.Printf("%x\n", x)
	t0 = time.Now()
	fmt.Println(PopCount(x))
	elapsed = time.Since(t0)
	fmt.Printf("%.8f s elapsed\n", elapsed.Seconds())
}
