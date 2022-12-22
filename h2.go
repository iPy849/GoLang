package main

import "fmt"

// “Create a function that concatenates two arrays into a new array.”

func concatArraysToArray(a [4]int, b [5]int) [9]int {
	// Esto es una clásico de primer semestre
	var c [len(a) + len(b)]int

	for i := 0; i < len(c); i++ {
		if i < len(a) {
			c[i] = a[i]
		} else {
			c[i] = b[i-len(a)]
		}
	}

	return c
}

func main() {
	a := [4]int{1, 2, 3, 4}
	b := [5]int{5, 6, 7, 8, 9}
	c := concatArraysToArray(a, b)
	fmt.Println(c)
}
