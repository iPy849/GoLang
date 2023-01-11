package main

import (
	"fmt"
)

func printer(ch chan<- bool, times int) {
	for i := 0; i < times; i++ {
		ch <- true
	}
	close(ch) // Cierra el canal
}

func main() {
	// This is an unbuffered channel
	var ch chan bool = make(chan bool)

	// Write 5 values to channel with a single goroutine and closes the channel
	go printer(ch, 5)

	// IMPORTANT: As the channel c is closed,
	// the range loop is going to exit on its own.
	for val := range ch {
		fmt.Print(val, " ")
	}
	fmt.Println()

	// This will always get the 0 value of the channel type
	for i := 0; i < 15; i++ {
		v, ok := <-ch
		fmt.Print(v, "-", ok, "\n")
	}
}
