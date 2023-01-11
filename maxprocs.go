package main

import (
	"fmt"
	"runtime"
)

/*
NOTE:: El paquete runtime tiene información sobre el sistema que
ejecuta el programa de go. Información como el compilador, la
arquitectura del procesador, versión de go, cantidad de procesadores
lógicos para goroutine scheduling máximos que se quiere usar
*/
func main() {
	fmt.Print("You are using ", runtime.Compiler, " ")
	fmt.Println("on a", runtime.GOARCH, "machine")
	fmt.Println("Using Go version", runtime.Version())
	// Esta línea setea la cantidad de procesadores lógicos a utilizar,
	// si esta se define menor a 1, entonces no cambia esta configuración,
	// tener en cuenta que este es un ajuste programático pero ya existe
	// una variable de entorno que contiene la información
	fmt.Printf("GOMAXPROCS: %d\n", runtime.GOMAXPROCS(0))
}
