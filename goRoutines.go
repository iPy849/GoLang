package main

import (
	"fmt"
	"time"
)

/*
Las gorrutinas pueden parecerse en concepto a un hilo/proceso/corrutina pero no lo son por lo que se
inventa el término gorrutina/goroutine

En pocas palabras va a mandar a ejecutar a una funcion y se va a desocupar de esta (parecido a lo que
hace una promesa en js). Según la documentación de Go, una gorrutina es: "it is a function executing
concurrently with other goroutines in the same address space. It is lightweight, costing little more
than the allocation of stack space".

Si bien trabajan sobre un mismo espacio de memoria, estos procesos se multiplexan sobre los procesos
del sistema operativo en si por lo que no es exactamente un proceso/hilo/corrutina pero si se apoya en
estos

Para llamar una gorrutina, tenemos que usar la sintaxis:
go <function call>(args...)

Las gorrutinas se inician en orden aleatorio y de eso se encarga el scheduler de Go
*/

func myPrint(start, finish int) {
	for i := start; i < finish; i++ {
		fmt.Print(i, " ")
	}
	fmt.Println()
	time.Sleep(100 * time.Microsecond)
}

func main() {
	for i := 0; i < 5; i++ {
		go myPrint(i, 5)
	}
	time.Sleep(time.Second)
}
