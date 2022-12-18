package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// os.Args siempre trae un slice/array con los argumentos: [binPath, arg1, ..., argn]
	// No importa si es un go build o si es un go run
	// NOTE: go run va a crear un binario y luego lo corre, cuando se acaba el programa, borra el binario

	// Comprueba cantidad de argumentos de cli
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Need one or more arguments")
		return
	}

	// Declara variables
	var min, max float64
	for i := 1; i < len(arguments); i++ {

		fmt.Printf("i=%d value -> %s \n", i, arguments[i])

		// Convierte string a float64
		n, err := strconv.ParseFloat(arguments[i], 64)
		if err != nil {
			continue
		}

		if n < min {
			min = n
		}
		if n > max {
			max = n
		}
	}
	fmt.Println("Min:", min)
	fmt.Println("Max:", max)
}
