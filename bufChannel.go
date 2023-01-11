package main

import (
	"fmt"
)

/*
	El objetivo de este script es demostrar que una buffered channel puede ser reutilizado
	y la llamada recursiva que añadí prueba que funciona, mientras no se a nil claro. Desde
	otro punto de vista esto funciona también como una cola y serviría para limitar la cantidad
	de tareas que se pueden gestionar
*/

var numbers = make(chan int, 5)
var repeat bool = true // NOTE: esta variable es mía

func main() {
	// numbers cannot store more than 5 integers
	counter := 10

	for i := 0; i < counter; i++ {
		select {
		// This is where the processing takes place
		case numbers <- i * i:
			fmt.Println("About to process", i)
		default:
			fmt.Println("No space for ", i)
		}
	}
	fmt.Println()

	for {
		select {
		case num := <-numbers:
			fmt.Println("*", num)
		default:
			fmt.Println("Nothing left to read!")

			// NOTE: este if es mío
			if repeat {
				repeat = false
				fmt.Println("Again----------")
				main()
			}
			return
		}
	}
}
