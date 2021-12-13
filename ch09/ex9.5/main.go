package main

import (
	"fmt"
	"time"
)

func main() {
	first := make(chan int)
	second := make(chan int)

	counter := 0
	elapsed := 0
	ticker := time.NewTicker(1 * time.Second)

	go func() {
		for {
			counter++
			second <- <-first
		}
	}()

	go func() {
		for {
			counter++
			first <- <-second
		}
	}()

	first <- 0
	for range ticker.C {
		elapsed++
		fmt.Printf("%d messages have been passed in %d seconds\n", counter, elapsed)
		if elapsed == 5 {
			close(first)
			close(second)
			break
		}
	}
}
