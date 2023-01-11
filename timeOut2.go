package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

/*
	NOTE: Este script es para hacer timeout desde una función externa pero la verdad no me
	gusta tanto como el primero donde se manajaba toda la lógica del timeout dentro de la misma
	función, no como en este caso que hay dos partes principales de la función, una con la lógica
	y una variable global con el resultado.

	Según yo, la entrada de cli simula el tiempo de la operación pero como que no me hace sentido
	script porque más bien prueba cuando no hay timeout, mi entrada tiene que ser mayor a 5s para
	que me de un resultado positivo y eso más bien indica lo contrario. Además usa la misma
	estrategia que timeOut1.go pero en una función fuera de main.
*/

// Variable global con el resultado del timeout
var result = make(chan bool)

func timeout(t time.Duration) {
	temp := make(chan int)
	// Se lanza una gorrutina que espera 5s (hardcodeado)
	go func() {
		time.Sleep(5 * time.Second)
		// Cuando termina de esperar, cierra el canal
		defer close(temp)
	}()

	// Inicia los listeners y espera a que uno se active
	select {
	case <-temp: // El canal se cerró con exito por lo que no ocurrió timeout
		result <- false
	case <-time.After(t): // Timeout
		result <- true
	}
}

func main() {
	arguments := os.Args
	if len(arguments) != 2 {
		fmt.Println("Please provide a time duration in milliseconds!")
		return
	}

	t, err := strconv.Atoi(arguments[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	duration := time.Duration(int32(t)) * time.Millisecond
	fmt.Printf("Timeout period is %s\n", duration)

	go timeout(duration)

	val := <-result
	if val {
		fmt.Println("Time out!")
	} else {
		fmt.Println("OK")
	}
}
