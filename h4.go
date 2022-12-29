package main

import "fmt"

type SA struct{}
type SB struct{}

func typeChecker(data interface{}) {
	switch v := data.(type) {
	case SA:
		fmt.Println("Heyyy data is SA")
	case SB:
		fmt.Println("Hello there, data is SB")
	default:
		fmt.Printf("Sorry... it is... an %T type?\n", v)
	}
}

func main() {
	typeChecker(SA{})
	typeChecker(SB{})
	typeChecker(33)
}
