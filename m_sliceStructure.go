package main

import "fmt"

type MyStruct struct {
	a int
	b int
	c bool
}

func main() {
	slc := make([]MyStruct, 4)
	fmt.Println("Structs in slice:", len(slc))
	for i, v := range slc {
		v.a = i
		v.b = i * i
		v.c = (i % 2) == 1
		fmt.Printf("%d. %v\n", i, v)
	}
}
