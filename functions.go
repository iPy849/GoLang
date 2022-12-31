package main

import "fmt"

func doubleSquare(x int) (int, int) {
	return x * 2, x * x
}

// Sorting from smaller to bigger value
func sortTwo(x, y int) (int, int) {
	if x > y {
		return y, x
	}
	return x, y
}

func main() {
	/*
		NOTE: Esto ya lo hicimos un montón de veces, no vale la pena
		explicarlo, lo único que merece ser mencionado es que devolver
		varias salidas simultáneamente nos ahorra escribir estructuras
		o dar la salida en tipos de colecciones
	*/
	n := 10
	d, s := doubleSquare(n)
	fmt.Println("Double of", n, "is", d)
	fmt.Println("Square of", n, "is", s)

	// NOTE: Una lambda salvaje ha aparecido!!!
	// An anonymous function
	anF := func(param int) int {
		return param * param
	}
	fmt.Println("anF of", n, "is", anF(n))

	fmt.Println(sortTwo(1, -3))
	fmt.Println(sortTwo(-1, 0))
}
