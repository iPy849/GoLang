package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {
	/*
		Números aleatorios en go:
			- Los sistemas digitales no generan números del aire, tienen procedimientos para ello
				y en realidad no llegan a existir números aleatorios sino que pseudoaleatorios
			- Se necesita de un seed/semilla para que go genere estos números
			- Hay una semilla default
			- No hay generación más que full aleatorio o aleatorios entre [0, n)
	*/
	seed := time.Now().Unix()
	rand.Seed(seed)

	for i := 0; i < 10; i++ {
		fmt.Println(strings.Repeat("-", 10))
		fmt.Println(rand.Intn(10))
		fmt.Println(rand.Intn(10))
	}
}
