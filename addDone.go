package main

import (
	"fmt"
	"os"
	"sync"
)

/*
NOTE: Aquí tenemos una situación interesante:
La cantidad de veces que que añadimos al contador no concuerda
con la cantidad de veces que le decimos que una gorrutina terminó
y por eso se va a quedar esperando indefinidamente cuando le digamos.
Esta condición se llama DeadLock (esperar indefinidamente por una acción que
no va a suceder).

También es muy común que cuando hacemos programación usando concurrencia
pase que no siempre el programa explote de la misma manera como en este caso
que explota 2/3 veces, eso lo hace más complicado de debugear
*/
func main() {
	count := 20
	fmt.Printf("Going to create %d goroutines.\n", count)

	flag := true
	if len(os.Args) == 1 {
		flag = false
	}

	var waitGroup sync.WaitGroup

	fmt.Printf("%#v\n", waitGroup)
	for i := 0; i < count; i++ {
		waitGroup.Add(1)
		go func(x int) {
			defer waitGroup.Done()
			fmt.Printf("%d ", x)
		}(i)
	}

	// More Add() calls
	if flag {
		waitGroup.Add(1)
	} else {
		// More Done() calls
		waitGroup.Done()
	}

	fmt.Printf("%#v\n", waitGroup)
	waitGroup.Wait()
	fmt.Println("\nExiting...")
}
