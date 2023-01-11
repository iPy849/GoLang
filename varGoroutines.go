package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
)

/*
NOTE: Aquí hay unas cuantas cositas importantes:
1- Usamos la libreria "sync"
2- Usamos la estructura sync.WaitGroup para llevar una cuenta de

	las gorrutinas en ejecución y esperarlas
*/
func main() {
	count := 10
	arguments := os.Args
	if len(arguments) == 2 {
		t, err := strconv.Atoi(arguments[1])
		if err == nil {
			count = t
		}
	}

	fmt.Printf("Going to create %d goroutines.\n", count)

	// Básicamente esto es un contador, la estructura tiene campos para llevar
	// la cuenta de las gorrutinas ejecutadas, tiene dos variables state que se
	// utilizan en dependencia de la arquitectura del sistema
	var waitGroup sync.WaitGroup
	fmt.Printf("%#v\n", waitGroup)
	for i := 0; i < count; i++ {
		// Al ser un contador tenemos que n cantidad de iniciaciones
		waitGroup.Add(1)
		go func(x int) {
			// Cuando se acaba la función, tenemos que indicar que ya terminamos
			// aunque en realidad esto es lo mismo que waitGroup.Add(-1)
			defer waitGroup.Done()
			fmt.Printf("%d ", x)
		}(i)
	}

	fmt.Printf("\n\n%#v\n", waitGroup)
	// Espera a que el contador state sea 0 para seguir con la siguiente instrucción
	waitGroup.Wait()
	fmt.Println("\nExiting...")
}
