package main

import (
	"fmt"
	"time"
)

func printme(x int) {
	fmt.Println("*", x)
	return
}

func main() {
	// NOTE: Crea una gorrutina para una función anónima
	go func(x int) {
		fmt.Printf("%d ", x)
	}(10)

	// NOTE: Crea una gorrutina para una función regular
	go printme(15)

	time.Sleep(time.Second / 2)
	fmt.Println("Exiting...")
}
