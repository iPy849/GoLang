package main

import (
	"fmt"
	"strings"
	"sync"
)

func writeToChannel(c chan int, x int) {
	c <- x
	close(c)
}

func printer(ch chan bool, i int) {
	fmt.Println(strings.Repeat("-", 5), " printer", i)
	ch <- true
}

func main() {
	//NOTE: Esto es un canal con buffer (tamaño 1 en este caso)
	c := make(chan int, 1)

	var waitGroup sync.WaitGroup
	waitGroup.Add(1)
	go func(c chan int) {
		defer waitGroup.Done()
		/*
			time.Sleep(time.Second * 5)

			en teoría cuando corra esto, me va a atrasar 5s la gorrutina
			y como es código concurrente no afecta a la gorrutina main, entonces
			TODO: ¿Por qué me está parando todo el programa antes de llegar a los prints que
			siguen o al switch?
		*/
		//NOTE: Escribe un 10 al canal y lo cierra
		writeToChannel(c, 10)
		fmt.Println("Exit.")
	}(c)

	// NOTE: Lee lo que hay en el canal
	fmt.Println("Read:", <-c)

	// NOTE: Comprueba que el canal está cerrado (igual que para saber si existe una llave en un mapa)
	_, ok := <-c
	if ok {
		fmt.Println("Channel is open!")
	} else {
		fmt.Println("Channel is closed!")
	}

	waitGroup.Wait()

	// NOTE: Esto es un canal sin buffer
	var ch chan bool = make(chan bool)
	// NOTE: 5 gorrutinas sin sincronización
	/*
		NOTE: Esto es una locura porque no siempre va a ejecutar estas 5
		gorrutinas antes de terminar el for que cierra el canal ch, entonces
		las gorrutinas que queden abiertas van a tratar de escribirle
		información a un canal cerrado y va a causar un panic
	*/
	for i := 0; i < 5; i++ {
		go printer(ch, i)
	}

	// Range on channels
	// IMPORTANT: As the channel c is not closed,
	// the range loop does not exit by its own.
	// NOTE: range funciona con las canales para saber si cerraron!!!
	n := 0
	for i := range ch {
		fmt.Println(i)
		if i == true {
			n++
		}
		if n > 2 {
			fmt.Println("n:", n)
			close(ch)
			break
		}
	}

	// NOTE: Aquí se lee 5 veces el canal en ch
	for i := 0; i < 5; i++ {
		fmt.Println(<-ch)
	}
}
