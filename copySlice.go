package main

import "fmt"

func main() {
	/*
		La función built-in copy(dst, src) copia los elementos que pueda del slice src
		al slice dst
	*/
	a1 := []int{1}
	a2 := []int{-1, -2}
	a5 := []int{10, 11, 12, 13, 14}
	fmt.Println("a1", a1)
	fmt.Println("a2", a2)
	fmt.Println("a5", a5)

	// copy(dst, src)
	// len(a2) > len(a1) so only -1 is copied into a1
	copy(a1, a2)
	fmt.Println("a1", a1)
	fmt.Println("a2", a2)

	// len(a5) > len(a1)
	copy(a1, a5)
	fmt.Println("a1", a1)
	fmt.Println("a5", a5)

	// len(a2) < len(a5) -> OK
	copy(a5, a2)
	fmt.Println("a2", a2)
	fmt.Println("a5", a5)
}
