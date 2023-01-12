package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)
/*
	NOTE: Este script tiene cambios en las entradas de las funciones add y send
	para que se viera que cuando es nil un canal, Go se paniquea feo y poder
	repartir el efecto.

	Este script muestra que GoLang se paniquea cuando tiene que leer o escribir
	de canales nil.
*/
var wg sync.WaitGroup

// Se cambió la entrada para que recibiera un apuntador, de otra manera el paso
// por valor solo copia el valor (al menos en Go)
func add(c *chan int) {
	sum := 0
	// NOTE: crea un timer y cuando esta listo manda un valor a un canal que tiene
	t := time.NewTimer(time.Second)

	for {
		select {
		case input := <-*c:
			sum = sum + input
		case <-t.C: // NOTE: Se envía el valor del timer por el canal que decía e indica que ya terminó
			fmt.Println(sum)
			// se hace nil el canal y al ser un apuntador todos los demás apuntadores van a recibir el
			// mismo resultado.
			*c = nil
			wg.Done()
			return
		}
	}
}

func send(c *chan int) {
	for {
		// Este if se cumple si le da tiempo a la gorrutina, normalmente no pasa
		if *c == nil{
			fmt.Println("Stopped")
			return
		} else {
			*c <- rand.Intn(10)
		}
	}
}

func main() {
	c := make(chan int)
	rand.Seed(time.Now().Unix())

	wg.Add(1)
	go add(&c)
	go send(&c)
	wg.Wait()
	// Aqui podemos ver que fue válido el nil que seteamos en add()
	if c == nil {
		fmt.Println("Soy nil")
	} else {
		fmt.Println("Existe")
	}
	// NOTE: esto se paniquea bien feo porque no se puede leer o escribir a canales nil
	/*
	fmt.Println(<-c)
	c->3
	 */
}
