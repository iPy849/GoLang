package main

import (
	"fmt"
	"math"
)

type Shape2D interface {
	Perimeter() float64
}

type circle struct {
	R float64
}

func (c circle) Perimeter() float64 {
	return 2 * math.Pi * c.R
}

func main() {
	// NOTE: Intenté usar type assertion en un struct pero lo único que conseguí
	// correr fue con reflections

	a := circle{R: 1.5}
	fmt.Printf("R %.2f -> Perimeter %.3f\n", a.R, a.Perimeter())

	// NOTE: Esta sintaxis es para type assertion sobre interfaces propias, va
	// a comprabar que "a" pueda satisfacer la interfaz Shape2D pero funciona
	// igual al type assertion que devuelve el valor y si es o no el tipo
	_, ok := interface{}(a).(Shape2D)
	if ok {
		fmt.Println("a is a Shape2D!")
	}

	i := 12
	_, ok = interface{}(i).(Shape2D)
	if ok {
		fmt.Println("i is a Shape2D!")
	}
}
