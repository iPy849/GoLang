package main

import "fmt"

/*
Go solo provee la palabra reservada "for" para ciclos (incluyendo while y foreach)
Se puede salir de un ciclo usando "break" y pasar una iteración con "continue"

A veces cuando una función te regresa varios valores, puedes omitir un valor con _, de la
misma forma que en python

// Like a C for
for init; condition; post { }

// Like a C while
for condition { }

// Like a C for(;;)
for { }

// References
https://go.dev/doc/effective_go#for
*/

func main() {
	// Traditional for loop
	for i := 0; i < 10; i++ {
		fmt.Print(i*i, " ")
	}
	fmt.Println()

	// While loop emulation
	i := 0
	for {
		if i == 10 {
			break
		}
		fmt.Print(i*i, " ")
		i++
	}
	fmt.Println()

	// For loop using range
	aSlice := []int{-1, 0, 1, 2, 3, 4}
	for i, v := range aSlice {
		if i%2 == 0 {
			continue
		}
		fmt.Println("index:", i, "value: ", v)
	}
}
