package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	/*
		Al igual que en python tenemos la función len para tamaños de colecciones
		y tenemos un paquete "os" que parece que permite interactuar con el os como dice su nombre
		y además devuelve un array
	*/
	if len(os.Args) != 2 {
		fmt.Println("Please provide a command line argument")
		return
	}
	argument := os.Args[1]

	/*
		Hay dos formas de escribir un switch:
		1: evaluando valores
		2: evaluando condiciones

		En Go tenemos la característica de que solo se ejecuta un case o el default sin
		necesidad de escribir un "break"
	*/
	switch argument {
	case "0":
		fmt.Println("Zero!")
	case "1":
		fmt.Println("One!")
	case "2", "3", "4":
		fmt.Println("2 or 3 or 4")
		fallthrough // Esta keyword sirve para saltar de este case al siguiente, que en este caso es default
	default:
		fmt.Println("Argument value:", argument)
	}

	// Atoi permite convertir strings a int y además te devuelve una respuesta de si fue posible la operación
	value, err := strconv.Atoi(argument)
	if err != nil {
		fmt.Println("Cannot convert to int:", argument)
		return
	}

	switch {
	case value == 0:
		fmt.Println("Zero!")
	case value > 0:
		fmt.Println("Positive Integer!")
	case value < 0:
		fmt.Println("Negative Integer!")
	default:
		fmt.Println("This should not happen:", value)
	}
}
