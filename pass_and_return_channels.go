package main

import (
	"fmt"
)

func main() {
	c := increment()
	cSum := get(c)
	for n := range cSum {
		fmt.Println(n)
	}
}

func increment() chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			out <- i
		}
		close(out)
	}()
	return out
}

func get(c chan int) chan int {
	out := make(chan int)
	go func() {
		var sum int
		for n := range c {
			sum += n
		}
		out <- sum
		close(out)
	}()
	return out
}
