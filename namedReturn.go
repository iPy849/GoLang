package main

import (
	"fmt"
	"os"
	"strconv"
)

/*
NOTE: Lo interesante de este script es que podemos definir
las variables de salida en la firma de la función
por lo que podemos usarlas en el cuerpo de la función
y cuando llamemos return va a devolver esas variables
en el orden que fueron declaradas, pero no estamos
limitados a ello pues también podemos devolver funciones
de manera regular
*/
func minMax(x, y int) (min, max int) {
	if x > y {
		return y, x
	}

	min = x
	max = y
	return
}

func main() {
	arguments := os.Args
	if len(arguments) < 3 {
		fmt.Println("The program needs at least 2 arguments!")
		return
	}

	// No checking here - we trust the user!!
	a1, _ := strconv.Atoi(arguments[1])
	a2, _ := strconv.Atoi(arguments[2])

	fmt.Println(minMax(a1, a2))
	mi, ma := minMax(a1, a2)
	fmt.Println(mi, ma)
}
