package main

import "fmt"

func suck(ch chan int) {
	fmt.Println(<-ch)
}

func main() {
	// without the capacity we will get a deadlock!
	cp := 20
	out := make(chan int, cp)
	out <- 5

	go suck(out)
}
