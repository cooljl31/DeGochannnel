package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)
	go sendData(ch)
	go getDate(ch)
	time.Sleep(1e9)
}

func getDate(ch chan string) {
	ch <- "Jim"
	ch <- "Maike"
	ch <- "Sirop"
	ch <- "Kaffe"
	ch <- "Kachu"
}

func sendData(ch chan string) {
	var input string
	for {
		input = <-ch
		fmt.Printf("%s ", input)
	}
}
