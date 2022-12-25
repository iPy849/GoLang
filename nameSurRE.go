package main

import (
	"fmt"
	"os"
	"regexp"
)

func matchNameSur(s string) bool {
	// Necesito hacer un slice de bytes para trabajar con regex
	t := []byte(s)
	// Comprueba que le Expresión regular sea correcta y devuelve un apuntador a un tipo *regexp.Regexp
	//	que tal pareciera ser una clase o similar
	re := regexp.MustCompile("^[A-Z][a-z]*$")
	fmt.Printf("\"re\" variable is type %T with value %v\n", re, re)
	// Este tipo de dato tiene adjuntos algunas funciones que permiten hacer el procesamiento con regex
	return re.Match(t)
}

func main() {
	// Recupera un argumento del cli
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Usage: <utility> string.")
		return
	}
	s := arguments[1]
	// Lo manda a la función matchNameSur para hacer todo el regex
	ret := matchNameSur(s)
	// Imprime el resultado
	fmt.Println(ret)
}
