package main

import "fmt"

// “Create a function that concatenates two arrays into a new slice.”

func concatArraysToSlice(a [4]int, b [5]int) []int {
	// Se puede usar [:] para obtener un slice de un array que abarca todos los elementos del array
	// Luego los anexamos a un nuevo slice y voila
	c := append([]int{}, a[:]...)
	c = append(c, b[:]...)
	return c
}

func main() {
	a := [4]int{1, 2, 3, 4}
	b := [5]int{5, 6, 7, 8, 9}
	c := concatArraysToSlice(a, b)
	fmt.Println(c)
}
