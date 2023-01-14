package main

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"time"
)

/*
NOTE: Como yo lo veo el contexto es un pseudo árbol para gorrutinas, el pseudo
árbol se empieza a expandir a partir de que se abre un contexto y cuando cierras
el contexto, vas a matar todas las gorrutinas, da igual si terminaron o no.

Hay varios tipos de contexto:
- Background (el contexto base)
	Este siempre tienes que matarlo con un cancel
- TimeOut
	Espera un tiempo y mata todo
- Deadline
	Espera un tiempo y mata todo --- TODO: investigar la diferencia
- withValue
	Básicamente tienes que igualar una propiedad a un valor que
	sea comparable, no se recomienda para nada usar los tipos
	built-in

NOTE: Hay que tener en cuenta que un contexto en realidad es una árbol
de contextos, donde un background pasa a ser timeout y mientras más lo
alargues, mayores propiedades alcanza.

NOTE: Para saber cuando acaba un contexto podemos usar un select
para escuchar un canal que todos los contexto tienen en su función
Done() y saber si terminó y luego puedes usar la función Err() para
tener un reporte de la resolución.

NOTE: El contexto siempre tiene que propagarse y por recomendación
y buenas prácticas siempre conviene pasarlo como primer argumento
*/

// The f1 function creates and executes a goroutine
// The time.Sleep() call simulates the time it would take a real goroutine
// to do its job - in this case it is 4 seconds. If the c1 context calls
// the Done() function in less than 4 seconds, the goroutine will not have
// enough time to finish.
func f1(t int) {
	c1 := context.Background()
	// WithCancel returns a copy of parent context with a new Done channel
	c1, cancel := context.WithCancel(c1)
	defer cancel()

	go func() {
		time.Sleep(4 * time.Second)
		cancel()
	}()

	select {
	case <-c1.Done():
		fmt.Println("f1() Done:", c1.Err())
		return
	case r := <-time.After(time.Duration(t) * time.Second):
		fmt.Println("f1():", r)
	}
	return
}

func f2(t int) {
	c2 := context.Background()
	c2, cancel := context.WithTimeout(c2, time.Duration(t)*time.Second)
	defer cancel()

	go func() {
		time.Sleep(4 * time.Second)
		cancel()
	}()

	select {
	case <-c2.Done():
		fmt.Println("f2() Done:", c2.Err())
		return
	case r := <-time.After(time.Duration(t) * time.Second):
		fmt.Println("f2():", r)
	}
	return
}

func f3(t int) {
	c3 := context.Background()
	deadline := time.Now().Add(time.Duration(2*t) * time.Second)
	c3, cancel := context.WithDeadline(c3, deadline)
	defer cancel()

	go func() {
		time.Sleep(4 * time.Second)
		cancel()
	}()

	select {
	case <-c3.Done():
		fmt.Println("f3() Done:", c3.Err())
		return
	case r := <-time.After(time.Duration(t) * time.Second):
		fmt.Println("f3():", r)
	}
	return
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Need a delay!")
		return
	}

	delay, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Delay:", delay)

	f1(delay)
	f2(delay)
	f3(delay)
}
