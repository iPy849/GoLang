package main

import "fmt"

func main() {
	a := [4]string{"zero", "one", "two", "three"}
	m := make(map[int]string)
	for i, v := range a {
		m[i] = v
	}
	fmt.Println(m)
}
