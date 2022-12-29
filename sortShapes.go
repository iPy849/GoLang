package main

import (
	"fmt"
	"math"
	"math/rand"
	"sort"
	"time"
)

const min = 1
const max = 5

// Esto es un generador aleatorio de decimales float64
func rF64(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

// Interfaz para Figuras 3D
// Going to sort the shapes based on their volume
type Shape3D interface {
	Vol() float64
}

// Estructuras
type Cube struct {
	x float64
}

type Cuboid struct {
	x float64
	y float64
	z float64
}

type Sphere struct {
	r float64
}

// Implementación del la función Vol() float64 del tipo
func (c Cube) Vol() float64 {
	return c.x * c.x * c.x
}

func (c Cuboid) Vol() float64 {
	return c.x * c.y * c.z
}

func (c Sphere) Vol() float64 {
	return 4 / 3 * math.Pi * c.r * c.r * c.r
}

// Slice of Shape3D
// Esto es un alias para decir para un slice de figuras 3d
type shapes []Shape3D

/*
NOTE: Aquí está lo interesante, este alias representa un tipo en sí mismo
por lo que se le pueden asociar funciones de tipo y es aquí que se implementan
las funciones de la interfaz sort.Interface para ordenar información
*/
// Implementing sort.Interface
func (a shapes) Len() int {
	return len(a)
}

func (a shapes) Less(i, j int) bool {
	return a[i].Vol() < a[j].Vol()
}

func (a shapes) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

// Esto es un print genérico para las diferentes figuras
func PrintShapes(a shapes) {
	for _, v := range a {
		// fmt.Printf("%.2f ", v)
		switch v.(type) {
		case Cube:
			fmt.Printf("Cube: volume %.2f\n", v.Vol())
		case Cuboid:
			fmt.Printf("Cuboid: volume %.2f\n", v.Vol())
		case Sphere:
			fmt.Printf("Sphere: volume %.2f\n", v.Vol())
		default:
			fmt.Println("Unknown data type!")
		}
	}
	fmt.Println()
}

func main() {
	data := shapes{}
	rand.Seed(time.Now().Unix())

	for i := 0; i < 3; i++ {
		cube := Cube{rF64(min, max)}
		cuboid := Cuboid{rF64(min, max), rF64(min, max), rF64(min, max)}
		sphere := Sphere{rF64(min, max)}

		data = append(data, cube)
		data = append(data, cuboid)
		data = append(data, sphere)
	}
	PrintShapes(data)

	// Sorting
	sort.Sort(shapes(data))
	PrintShapes(data)

	// Reverse sorting
	sort.Sort(sort.Reverse(shapes(data)))
	PrintShapes(data)
}
