package main

import "fmt"

/*
Se supone que any es un alias de una interfaz pero de esto se habla en más profundidad en
el capítulo 4 y 13
*/
func Print[T any](s []T) {
	for _, v := range s {
		fmt.Print(v, " ")
	}
	fmt.Println()
}

func main() {
	Ints := []int{10, 22, 3}
	Strings := []string{"01", "2", "33"}
	Print(Ints)
	Print(Strings)
}
