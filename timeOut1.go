package main

import (
	"fmt"
	"time"
)

/*
Aquí pasa que se simula una operación de timeout donde la
operación a realizar se tarda más de lo esperado
*/
func main() {
	// Crea canal y corre gorrutina lambda que necesita 3s pero solo espera 1s
	c1 := make(chan string)
	go func() {
		time.Sleep(3 * time.Second)
		c1 <- "c1 OK"
	}()

	select {
	// NOTE: Asigna el valor y es listener a la vez
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(time.Second):
		fmt.Println("timeout c1")
	}

	// Crea canal y corre gorrutina lambda que necesita 3s pero solo espera 4s y se completa
	c2 := make(chan string)
	go func() {
		time.Sleep(3 * time.Second)
		c2 <- "c2 OK"
	}()

	select {
	case res := <-c2:
		fmt.Println(res)
	case <-time.After(4 * time.Second):
		fmt.Println("timeout c2")
	}
}
