package main

import (
	"fmt"
	"strconv"
	"sync"
)

var wg sync.WaitGroup

func main() {
	SharedResource := "Default text"
	var mutex sync.RWMutex

	for i := 0; i < 2; i++ {
		wg.Add(3)

		go func(i int) {
			fmt.Println("Lanza lectura", i)
			defer wg.Done()
			mutex.RLock()
			fmt.Println("R Lock", i)
			// time.Sleep(time.Second * 2)
			fmt.Println("Lectura:", SharedResource)
			mutex.RUnlock()
			fmt.Println("R Unlock", i)
		}(i)

		go func(i int) {
			fmt.Println("Lanza escritura", i)
			defer wg.Done()
			// time.Sleep(time.Second)
			mutex.Lock()
			fmt.Println("W Lock", i)
			SharedResource = SharedResource + strconv.Itoa(i)
			mutex.Unlock()
			fmt.Println("W Unlock", i)
		}(i)

		go func(i int) {
			fmt.Println("Lanza segunda lectura", i)
			defer wg.Done()
			mutex.RLock()
			fmt.Println("R 2 Lock", i)
			// time.Sleep(time.Second * 3)
			fmt.Println("Segunda lectura:", SharedResource)
			mutex.RUnlock()
			fmt.Println("R 2 Unlock", i)
		}(i)
	}
	wg.Wait()
	fmt.Println("Finalmente:", SharedResource)

}
