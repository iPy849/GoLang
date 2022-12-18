package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Need one or more arguments")
		return
	}

	/*
		Se supone que esto es un técnica para no cometer errores de entrada en los programas
		El prinicipio es recorrer todos los argumentos de cli y ver cúantos se pueden leer, y de
		estos cuántos son enteros y cuántos son flotantes
	*/

	var total, nInts, nFloats int
	invalid := make([]string, 0)
	for _, k := range arguments[1:] {
		_, err := strconv.Atoi(k)
		if err == nil {
			total++
			nInts++
			continue
		}

		_, err = strconv.ParseFloat(k, 64)
		if err == nil {
			total++
			nFloats++
			continue
		}

		invalid = append(invalid, k)
	}
	fmt.Printf("Read: %d Ints: %d Floats: %d\n", total, nInts, nFloats)
}
