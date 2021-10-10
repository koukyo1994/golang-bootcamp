package main

import "fmt"

// pc[i]はiのポピュレーションカウント
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCountはxのポピュレーションカウント(1が設定されているビット数)を返します
func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

func PopCountLoop(x uint64) int {
	var count int
	for i := 0; i < 8; i++ {
		idx := byte(x >> (i * 8))
		count += int(pc[idx])
	}
	return count
}

func PopCountShift(x uint64) int {
	var count int
	for i := 0; i < 64; i++ {
		if x&1 == 1 {
			count++
		}
		x >>= 1
	}
	return count
}

func PopCountClear(x uint64) int {
	var count int
	for x != 0 {
		x &= x - 1
		count++
	}
	return count
}

func main() {
	var x uint64 = 0x1234567890ABCDEF
	fmt.Printf("%x\n", x)
	fmt.Println(PopCount(x))
	fmt.Println(PopCountLoop(x))
	fmt.Println(PopCountShift(x))
	fmt.Println(PopCountClear(x))
}
