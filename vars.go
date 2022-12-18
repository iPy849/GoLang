package main

import (
	"fmt"
	"math"
)

/*
	Toda variable o función que no sea local, debe empezar con la palabra reservada
	var, const o func y creo que la sintaxis se entiende bastante bien
*/
var Global int = 1234
var AnotherGlobal = -5678

func main() {
	// Al declarar una variable y no asignarle un valor inicial, se inicializa en su valor default
	var j int
	// Esto se llama: short assignment statement y permite inicializar una variable, asignarle un valor e inferir el tipo
	i := Global + AnotherGlobal
	fmt.Println("Initial j value:", j)
	j = Global
	/*
		Go no hace cambios de tipos implícitamente, hay que hacer estos cambios de manera manual como en
		el siguiente caso
	*/
	k := math.Abs(float64(AnotherGlobal))
	fmt.Printf("Global=%d, i=%d, j=%d k=%.2f.\n", Global, i, j, k)
}
