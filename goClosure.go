package main

import (
	"fmt"
	"sync"
)

/*
NOTE: Una closure evalúa todas las variables externas cuando se ejecuta
por primera vez
*/

var w sync.WaitGroup

func main() {
	for i := 0; i <= 20; i++ {
		fmt.Println("Itero", i)
		w.Add(1)
		// NOTE: Hola soy una closure
		go func() {
			fmt.Println("Lanzo", i)
			defer w.Done()
			fmt.Println("...", i)
		}()
	}
	w.Wait()
	fmt.Println()
}
