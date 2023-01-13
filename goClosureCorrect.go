package main

import (
	"fmt"
	"time"
)

func main() {

	// NOTE: Una forma de corregir una closure y los valores que tiene
	for i := 0; i <= 20; i++ {
		i := i
		go func() {
			fmt.Print(i, " ")
		}()
	}
	time.Sleep(time.Second)
	fmt.Println()

	// NOTE: La forma que me gusta de corregir una closure
	for i := 0; i <= 20; i++ {
		go func(x int) {
			fmt.Print(x, " ")
		}(i)
	}
	time.Sleep(time.Second)
	fmt.Println()
}
