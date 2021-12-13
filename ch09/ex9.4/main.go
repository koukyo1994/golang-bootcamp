package main

import (
	"flag"
	"fmt"
	"time"
)

var n = flag.Int("n", 2, "number of steps of the pipeline")

func main() {
	flag.Parse()
	channels := make([]chan int, *n)
	for i := 0; i < *n; i++ {
		channels[i] = make(chan int)
	}

	t0 := time.Now()
	// First method
	go func() {
		for x := 0; x < 100; x++ {
			channels[0] <- x
		}
		close(channels[0])
	}()

	for i := 1; i < *n; i++ {
		go func(i int) {
			for x := range channels[i-1] {
				channels[i] <- x
			}
			close(channels[i])
		}(i)
	}

	// Main
	for x := range channels[*n-1] {
		fmt.Println(x)
	}
	fmt.Printf("Elapsed: %.2f s \n", time.Since(t0).Seconds())
}
