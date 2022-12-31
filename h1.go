/*
	Está el paquete sort pero el que haga un burbujas es a propósito de
	nostalgia
*/

package main

import "fmt"

func sortThreeValsV1(a, b, c int) [3]int {
	v := [3]int{a, b, c}

	for i := 0; i < 3; i++ {
		for j := i; j < 3; j++ {
			if v[i] < v[j] {
				tmp := v[j]
				v[j] = v[i]
				v[i] = tmp
			}
		}
	}
	return v
}

func sortThreeValsV2(a, b, c int) (v [3]int) {
	v = [3]int{a, b, c}

	for i := 0; i < 3; i++ {
		for j := i; j < 3; j++ {
			if v[i] < v[j] {
				tmp := v[j]
				v[j] = v[i]
				v[i] = tmp
			}
		}
	}
	return
}

func main() {
	// Me gusta más la versión 1 porque me resulta más familiar
	fmt.Println(sortThreeValsV1(3, 1, 6))
	fmt.Println(sortThreeValsV2(3, 1, 6))
}
