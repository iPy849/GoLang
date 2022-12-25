package main

import (
	"fmt"
	"os"
)

type Name struct {
	first string
	last  string
	index int
}

func main() {
	// Filra las el formato de los argumentos
	al := len(os.Args)
	if al == 1 || al%2 == 0 {
		fmt.Println("usage: h3.go ...[<name> <last name>]")
		return
	}

	// Extrae los argumentos y crea slice de Name
	a := os.Args[1:]
	sName := make([]Name, 0)

	// Itera cada pareja de argumentos y lo agrega a la salida como un estructura
	for i := 0; i < len(a)/2; i++ {
		name := a[2*i]
		lastName := a[(2*i)+1]
		tmpName := Name{name, lastName, i}
		sName = append(sName, tmpName)
	}
	// Salida
	fmt.Println(sName)
}
