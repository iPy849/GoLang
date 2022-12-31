package main

import "fmt"

/*
	NOTE: El objetivo de esta sección es devolver funciones
	y su sintaxis como tipo de salida. La verdad se entiende
	con solo mirarlo
*/

func funRet(i int) func(int) int {
	if i < 0 {
		return func(k int) int {
			k = -k
			return k + k
		}
	}

	return func(k int) int {
		return k * k
	}
}

func main() {
	n := 10
	i := funRet(n)
	// The -4 parameter is used for determining
	// the anonymous function that will be returned
	j := funRet(-4)

	fmt.Printf("%T\n", i)
	fmt.Printf("%T %v\n", j, j)
	fmt.Println("j", j, j(-5))

	// Same input parameter but DIFFERENT
	// anonymous functions assigned to i and j
	fmt.Println(i(10))
	fmt.Println(j(10))
}
