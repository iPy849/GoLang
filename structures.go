package main

import "fmt"

// Knowning the data structures of a program is really important
// Programs = Data Structures + Algorithms
type Entry struct {
	Name    string
	Surname string
	Year    int
}

// Initialized by Go
func zeroS() Entry {
	return Entry{}
}

// Initialized by the user
func initS(N, S string, Y int) Entry {
	if Y < 2000 {
		return Entry{Name: N, Surname: S, Year: 2000}
	}
	return Entry{Name: N, Surname: S, Year: Y}
}

// Initialized by Go - returns pointer
func zeroPtoS() *Entry {
	t := &Entry{}
	return t
}

// Initialized by the user - returns pointer
func initPtoS(N, S string, Y int) *Entry {
	if len(S) == 0 {
		return &Entry{Name: N, Surname: "Unknown", Year: Y}
	}
	return &Entry{Name: N, Surname: S, Year: Y}
}

func main() {
	// Crea una structura entry
	s1 := zeroS()
	// Crea una structura entry y devuelve un puntero a la estructura
	p1 := zeroPtoS()
	// Aquí se imprime la estructrua y dereferencia el puntero y se imprime
	fmt.Println("s1:", s1, "p1:", *p1)

	// Crea una structura entry
	s2 := initS("Mihalis", "Tsoukalos", 2020)
	// Crea una structura entry y devuelve un puntero a la estructura
	p2 := initPtoS("Mihalis", "Tsoukalos", 2020)
	// Aquí se imprime la estructrua y dereferencia el puntero y se imprime
	fmt.Println("s2:", s2, "p2:", *p2)
	fmt.Println("Year:", s1.Year, s2.Year, p1.Year, p2.Year)

	// Se inicializa una instancia de entry a partir de new() que regresa un puntero a un entry vacío con todos
	// sus campos en valor cero
	pS := new(Entry)
	fmt.Println("pS:", pS)
}
