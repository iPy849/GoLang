package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	/*
		Go no tiene forma de borrar elementos de un slice, por eso tenemos que implemntarlo manualmente
		pero también explica dos técnicas.

		# La primera me gusta más porque se hace en un sola línea
		1. Sacar dos slices del slice al que borrar elementos, los dos slices excluyen el elemento target
			y finalmente se unen los dos slices y lo almacena en la variable original

		2. Copiar el último elemento del slice al que borrar elementos en el target y devolver un slice
			que excluya el último dato y almacenarlo en la variable
	*/

	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Need an integer value.")
		return
	}

	index := arguments[1]
	i, err := strconv.Atoi(index)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Using index", i)

	aSlice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println("Original slice:", aSlice)

	// Delete element at index i
	if i > len(aSlice)-1 {
		fmt.Println("Cannot delete element", i)
		return
	}

	// The ... operator auto expands aSlice[i+1:] so that
	// its elements can be appended to aSlice[:i] one by one
	aSlice = append(aSlice[:i], aSlice[i+1:]...)
	fmt.Println("After 1st deletion:", aSlice)

	// Delete element at index i
	if i > len(aSlice)-1 {
		fmt.Println("Cannot delete element", i)
		return
	}

	// Replace element at index i with last element
	aSlice[i] = aSlice[len(aSlice)-1]
	// Remove last element
	aSlice = aSlice[:len(aSlice)-1]
	fmt.Println("After 2nd deletion:", aSlice)
}