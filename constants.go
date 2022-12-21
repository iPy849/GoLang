package main

import (
	"fmt"
)

/*
La sintaxis:

type <Name> <Underlying type>

permite declarar un tipo nuevo que su valor es de otro tipo (especificafo).
Es algo parecido (no igual) a lo que pasa con las runas y los int32
*/
type Digit int
type Power2 int

const PI = 3.1415926

const (
	C1 = "C1C1C1"
	C2 = "C2C2C2"
	C3 = "C3C3C3"
)

func main() {
	const s1 = 123
	var v1 float32 = s1 * 12
	fmt.Println(v1)
	fmt.Println(PI)

	/*
		La palabra reservada iota es un generador autoincremental que se reinicia
		cuando se define otra constante, un ejemplo claro de lo que hace a continuación
	*/
	const (
		Zero Digit = iota
		One
		Two
		Three
		Four
	)

	fmt.Println(One)
	fmt.Println(Two)

	/*
		Aquí hay dos cosas interesantes:
		1- Tenemos un _ en una constante y aparte con un iota, esto permite saltarse un valor
			de los que genera iota
		2- Tenemos al operador << que es una corrida de bits para el 1, es decir, corre 00000001 iota veces
	*/
	const (
		p2_0 Power2 = 1 << iota
		_
		p2_2
		_
		p2_4
		_
		p2_6
	)

	fmt.Println("2^0:", p2_0)
	fmt.Println("2^2:", p2_2)
	fmt.Println("2^4:", p2_4)
	fmt.Println("2^6:", p2_6)
}
