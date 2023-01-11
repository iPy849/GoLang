package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"sync"
	"time"
)

func gen(min, max int, createNumber chan int, end chan bool) {
	// time.Sleep(time.Second) // NOTE: La verdad está aquí de más
	for { // NOTE: El select hay que meterlo en un ciclo infinito
		select {
		// case createNumber <- rand.Intn(max-min) + min:
		case <-end:
			fmt.Println("Ended!")
			return
		case <-time.After(3 * time.Second):
			fmt.Println("time.After()!")
			return
		}
	}
}

func main() {
	rand.Seed(time.Now().Unix())
	createNumber := make(chan int)
	end := make(chan bool)
	fmt.Println(&end)

	if len(os.Args) != 2 {
		fmt.Println("Please give me an integer!")
		return
	}

	n, _ := strconv.Atoi(os.Args[1])
	fmt.Printf("Going to create %d random numbers.\n", n)

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		gen(0, 2*n, createNumber, end)
		wg.Done()
	}()

	// NOTE:
	go func() {
		for i := 0; i < n; i++ {
			v, ok := <-createNumber
			if ok {
				fmt.Printf("%d - %d \n", i, v)
			}
		}
		end <- true
		wg.Done()
	}()

	wg.Wait()
	close(createNumber)
	close(end)
	fmt.Println(&end)
	fmt.Println("Exiting...")
}
