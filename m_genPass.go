package main

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

// Esta función devuelve un número entre [min, max)
func randBtw(min int, max int) (int, error) {
	if min > max {
		return 0, errors.New("func randBtw(min int, max int) -> error: max < min")
	}
	return rand.Intn(max-min) + min, nil
}

func main() {
	var passLength int

	// Control de argumentos de cli
	if len(os.Args) == 2 {
		tmpV := os.Args[1]
		length, err := strconv.Atoi(tmpV)

		if err != nil || length <= 0 {
			fmt.Println("Setted default length")
			length = 8
		}
		passLength = length
	} else {
		fmt.Println("Usage: m_genPass.go MAX_LENGTH")
	}

	// Es más fácil tener el ascii desde una runa,
	startCharacter := int('!')
	pass := "" // Resultado de la generación
	// Semilla de generación numérica
	rand.Seed(time.Now().Unix())
	for i := 0; i < passLength; i++ {
		// Genera el número
		ranNumber, err := randBtw(startCharacter, 94)
		if err != nil {
			fmt.Println(err)
			os.Exit(-1)
		}
		// Convierte la representación ASCII del número en string y lo concatena
		pass = pass + string(ranNumber)
	}
	println(pass)
}
