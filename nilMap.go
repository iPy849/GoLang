package main

import (
	"fmt"
)

func main() {
	// Inicializa mapa vacío y un valor
	aMap := map[string]int{}
	aMap["test"] = 1
	fmt.Println("aMap:", aMap)
	// Asiga nil a la variable que tenía referenciado al mapa
	aMap = nil
	fmt.Println("aMap:", aMap)
	if aMap == nil {
		fmt.Println("nil map!")
		// Asigna otro mapa
		aMap = map[string]int{}
	}
	aMap["test"] = 1

	// This will crash!
	aMap = nil
	// Por tratar de acceder un mapa nil
	aMap["test"] = 1
}
