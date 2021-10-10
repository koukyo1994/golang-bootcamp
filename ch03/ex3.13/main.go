package main

import (
	"fmt"
)

const (
	KB = 1000
	MB = KB * KB
	GB = MB * KB
	TB = GB * KB
	PB = TB * KB
	EB = PB * KB
	ZB = EB * KB
	YB = ZB * KB
)

func main() {
	fmt.Printf("1KB = %d bytes\n", KB)
	fmt.Printf("1MB = %d bytes\n", MB)
	fmt.Printf("1GB = %d bytes\n", GB)
	fmt.Printf("1TB = %d bytes\n", TB)
	fmt.Printf("1PB = %d bytes\n", PB)
	fmt.Printf("1EB = %d bytes\n", EB)
	// fmt.Printf("1ZB = %s bytes", ZB)
	// fmt.Printf("1YB = %s bytes", YB)
}
