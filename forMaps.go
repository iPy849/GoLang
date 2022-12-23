package main

import "fmt"

func main() {
	// range works with maps as well
	aMap := make(map[string]string)
	aMap["123"] = "456"
	aMap["key"] = "A value"
	// Itera obteniendo ambos valores (índice, valor) por registro
	for key, v := range aMap {
		fmt.Println("key:", key, "value:", v)
	}

	// Como se dijo antes, se puede saltar el recibir un valor con _
	for _, v := range aMap {
		fmt.Print(" # ", v)
	}
	fmt.Println()
}
