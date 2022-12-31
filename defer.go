package main

import (
	"fmt"
)

/*
NOTE: Aquí tenemos que d2() va a imprimr el valor 0 de i ¿Por qué?
Aquí pasa que como el concepto de lambdas y closure es intercambiable en
Go, las funciones anónimas tienen acceso al scope donde se definen pero,
si desaparece la variabla referenciada, va a tomar el valor 0 del tipo.
En cambio, si pasas un valor como parámetro, ese parámetro forma parte del
scope de la función y no desparece hasta que la función concluye, justo lo
que pasa con d1() y d3()
*/

func d1() {
	for i := 3; i > 0; i-- {
		defer fmt.Print(i, " ")
	}
}

func d2() {
	for i := 3; i > 0; i-- {
		defer func() {
			fmt.Print(i, " ")
		}()
	}
	fmt.Println()
}

func d3() {
	for i := 3; i > 0; i-- {
		defer func(n int) {
			fmt.Print(n, " ")
		}(i)
	}
}

func main() {
	d1()
	d2()
	fmt.Println()
	d3()
	fmt.Println()
}
