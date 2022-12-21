package main

import "fmt"

func main() {
	aString := "Hello world!!! €"
	fmt.Println("First character:", aString[0])

	//Runes
	// Asi se declara una runa
	/*
		Una runa es una representación de un caracter unicode, el dato que representa
		es un int32 pero su tipo es runa. Si intentas representar una runa de otra modo
		que no sea como caracter %c va a devolver su representación int32
	*/
	r := '€'
	fmt.Println("As an int32 value:", r)

	// Convert Runes to text
	fmt.Printf("As a string: %s and as a character: %c\n", r, r)

	// Print an existing string as runes
	for _, v := range aString {
		fmt.Printf("%x ", v)
	}
	fmt.Println()
}
