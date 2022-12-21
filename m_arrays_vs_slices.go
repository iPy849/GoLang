package main

import (
	"fmt"
	"strings"
)

const SEPARATOR_LENGTH int = 20

func main() {
	/*
		Slices vs arrays basics:

		Arrays:
		1. Tamaño estático después de la definición
		2. Se tiene que especificar el tamaño de un array
		3. Cuando se pasa a una función, se pasa una copia del array, no el original

		Slice:
		1. Es una estructura dinámica que se basa en arrays pero es otro tipo
		2. Tamaño variable
		3. Realmente lo que se pasa son una estructura que representa un apuntador a
			la cabecera del array por lo que los cambios se hacen sobre el array original
			no sobre una copia
	*/
	a := make([]int, 3)
	fmt.Printf("%v %T\n", a, a)
	a = append(a, 1)
	fmt.Printf("%v %T\n", a, a)
	doingChanges2Slice(a, 1, 2)
	fmt.Printf("After function %v %T\n", a, a)
	fmt.Println(strings.Repeat("-", SEPARATOR_LENGTH))

	b := [4]int{}
	fmt.Printf("%v %T\n", b, b)
	b[3] = 1
	fmt.Printf("%v %T\n", b, b)
	doingChanges2Array(b, 1, 3)
	fmt.Printf("After function %v %T\n", b, b)
	fmt.Println(strings.Repeat("-", SEPARATOR_LENGTH))
}

func doingChanges2Array(arr [4]int, index int, v int) [4]int {
	arr[index] = v
	fmt.Println("doing changes to array", arr, "inside the function")
	return arr
}

func doingChanges2Slice(slice []int, index int, v int) []int {
	slice[index] = v
	fmt.Println("doing changes to slice", slice, "inside the function")
	return slice
}
