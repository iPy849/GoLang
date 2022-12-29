package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

// Estructura definida
type MyStructure struct {
	elem1 int
	elem2 string
}

// Slice de estructura
type MyStructureCollection []MyStructure

func (a MyStructureCollection) Len() int {
	return len(a)
}

func (a MyStructureCollection) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a MyStructureCollection) Less(i, j int) bool {
	if a[i].elem1 == a[j].elem1 {
		return a[i].elem2 < a[j].elem2
	}
	return a[i].elem1 < a[j].elem1
}

// Genera números aleatorios
func random(min, max int) int {
	return min + rand.Intn(max-min)
}

// Genera strings aleatorios
func generateRandomString(n int) string {
	var r string
	for i := 0; i < n; i++ {
		r = r + string(random(35, 125))
	}
	return r
}

func main() {
	// Genera los datos
	var collection MyStructureCollection
	rand.Seed(time.Now().Unix())
	for i := 0; i < 10; i++ {
		data := MyStructure{
			random(0, 100),
			generateRandomString(10),
		}
		collection = append(collection, data)
	}
	fmt.Println(collection)

	// Ordena la información (Otra vez)
	sort.Sort(collection)
	fmt.Println(collection)

	sort.Sort(sort.Reverse(collection))
	fmt.Println(collection)
}
