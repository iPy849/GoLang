package main

import "fmt"

func main() {
	m := map[int]string{0: "zero", 1: "one", 2: "two", 3: "three"}
	ks := []int{}
	vs := []string{}
	for i, v := range m {
		ks = append(ks, i)
		vs = append(vs, v)
	}
	fmt.Printf("index: %v\nvalues:%v\n", ks, vs)
}
