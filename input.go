package main

import "fmt"

func main() {
	fmt.Printf("Say my name: ")
	var name string
	// TODO: Necesito saber una forma de que una función exprese que quiere un puntero y no un valor
	// Se explica el tema de punteros en capítulos posteriores
	fmt.Scanln(&name)
	fmt.Println("My name is", name)
}
