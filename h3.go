package main

import "fmt"

// “Create a function that concatenates two slices into a new array.”

// TODO: Preguntar esto a alguien que sepa de go en enero
func concatSliceToArray(a []int, b []int) [9]int {
	var c [9]int

	for i := 0; i < len(a)+len(b); i++ {
		if i < len(a) {
			c[i] = a[i]
		} else {
			c[i] = b[i-len(a)]
		}
	}

	return c
}

func main() {
	a := []int{1, 2, 3, 4}
	b := []int{5, 6, 7, 8, 9}
	c := concatSliceToArray(a, b)
	fmt.Println(c)
}
