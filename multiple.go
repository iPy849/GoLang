package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

/*
NOTE: Nada especial pasa acá, leyendo el código se entiende
que levanta n gorrutinas como le especifiques en el cli.

Un detalle es que todavía está usando time.Sleep para esperar
y que no se mueran todas las gorrutinas levantadas
*/
func main() {
	count := 10
	if len(os.Args) == 1 {
		fmt.Println("Using default value.")
	} else {
		temp, err := strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Println("Using default value.")
		} else {
			count = temp
		}
	}

	fmt.Printf("Going to create %d goroutines.\n", count)
	for i := 0; i < count; i++ {
		go func(x int) {
			fmt.Printf("%d ", x)
		}(i)
	}
	time.Sleep(time.Second)
	fmt.Println("\nExiting...")
}
