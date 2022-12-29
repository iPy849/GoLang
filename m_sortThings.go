package main

import (
	"fmt"
	"reflect"
	"sort"
)

// Implementación de typo que cumpla la interfaz Sort
type Animals []Animal

func (animals Animals) Len() int {
	return len(animals)
}

func (animals Animals) Less(i, j int) bool {
	/*
		Devuelve la comparación donde i < j

		Pero yo quiero comparar por nombre pero lo que traigo es una interfaz
		y no tengo acceso a priori al tipo de la estructura o sus campos,
		por eso uso  dos estrategias:
		1- el paquete reflect para comprobar el nombre del tipo
			y hacer type assertion sin que se paniquee el programa y
			así llegar al valor
		2- un switch type para comprobar el tipo y recuperar el valor
	*/
	names := [2]string{}

	if reflect.TypeOf(animals[i]).String() == "main.LandAnimal" {
		names[0] = animals[i].(LandAnimal).name
	} else {
		names[0] = animals[i].(WaterAnimal).name
	}

	switch v := animals[j].(type) {
	case LandAnimal:
		names[1] = v.name
	case WaterAnimal:
		names[1] = v.name
	}

	return names[0] < names[1]
}

func (animals Animals) Swap(i, j int) {
	animals[i], animals[j] = animals[j], animals[i]
}

// Definición de la intefaz Animal
type Animal interface {
	doAnything()
}

// Definición de estructura LandAnimal y funciones para interfaz Animal
type LandAnimal struct {
	name string
	legs int
}

func (la LandAnimal) doAnything() {
	fmt.Printf("%s walks and runs using his %d legs\n", la.name, la.legs)
}

// Definición de estructura WaterAnimal y funciones para interfaz Animal
type WaterAnimal struct {
	name                string
	canBreathOutOfWater bool
}

func (wa WaterAnimal) doAnything() {
	var canBreathText string
	if wa.canBreathOutOfWater {
		canBreathText = "can"
	} else {
		canBreathText = "can't"
	}
	fmt.Printf("%s can swim but he %s breath out of water\n", wa.name, canBreathText)
}

func main() {
	/*
		La diferencia entre este script y sortShapes.go es que en este
		uso la interfaz para comportamiento de los datos pero cuando quiero
		implementar la interfaz sort voy a comparar por nombre.

		Dejo comentarios por todo el código
	*/

	// Datos hardcodeados
	data := Animals{
		LandAnimal{"Caballo", 4},
		LandAnimal{"Humano", 2},
		WaterAnimal{"Pez", false},
		WaterAnimal{"Ascua", true},
	}
	fmt.Println(data)

	// Ya que Animals es un alias para []Animal puedo tomar cada
	// elemento como que tiene tipi Animal
	for _, v := range data {
		fmt.Println(v)
		v.doAnything() // Puedo llamar el método de la interfaz
	}

	// Organiza la información usando la interfaz Sort
	sort.Sort(data)
	fmt.Println(data)

	sort.Sort(sort.Reverse(data))
	fmt.Println(data)

}
