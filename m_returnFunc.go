package main

import "fmt"

// NOTE: Solo devuelve funciones distintas
func SaySomething(areYouOk bool) func() {
	if areYouOk {
		return func() {
			fmt.Println("Hola hermoso mundo!!!")
		}
	}
	return func() {
		fmt.Println("Hola mundo cruel ...")
	}
}

func main() {
	// Almacena cada función anónima
	f1 := SaySomething(true)
	f2 := SaySomething(false)
	// Ejecuta las funciones
	f1()
	f2()
	// Detalles
	fmt.Printf("f1 -> val: %v | type: %T \n", f1, f1)
	fmt.Printf("f2 -> val: %v | type: %T \n", f2, f2)
}
