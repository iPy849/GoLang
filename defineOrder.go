package main

import (
	"fmt"
	"sync"
	"time"
)

/*
La idea de este script es ejecutar gorrutinas en un orden específico.
La técnica solo funciona con pocas gorrutinas y para ello va a tener
que paralizar cada una con una instrucción de recibir de un canal y cuando
termina esa gorrutina cierra el canal que recibe y el próximo canal para que
la próxima función reciba un valor 0 y pueda ejecutarse. De esa manera llegamos
hasta la función D que se llama cuatro veces en gorrutinas distintas, cuando
se libera por el canal se ejecutan las 4 de forma concurrente.
*/

var wg sync.WaitGroup

func A(a, b chan struct{}) {
	<-a
	fmt.Println("A()!")
	time.Sleep(time.Second)
	close(b)
}

func B(a, b chan struct{}) {
	<-a
	fmt.Println("B()!")
	time.Sleep(3 * time.Second)
	close(b)
}

func C(a, b chan struct{}) {
	<-a
	fmt.Println("C()!")
	close(b)
}

func D(a chan struct{}) {
	<-a
	fmt.Println("D()!")
	wg.Done()
}

func main() {
	x := make(chan struct{})
	y := make(chan struct{})
	z := make(chan struct{})
	w := make(chan struct{})

	wg.Add(1)
	go func() {
		D(w)
	}()

	wg.Add(1)
	go func() {
		D(w)
	}()

	go A(x, y)

	wg.Add(1)
	go func() {
		D(w)
	}()

	go C(z, w)
	go B(y, z)

	wg.Add(1)
	go func() {
		D(w)
	}()

	// This triggers the process
	close(x)
	wg.Wait()
}
