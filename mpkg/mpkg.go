package mpkg

import "fmt"

func init() {
	fmt.Println("Hola soy el init de mpkg package")
}

var YELL string = "Heyyy look, I'm global"

func SayHello() {
	fmt.Println("Hello")
}

var something string = "reading something..."

func ReadmeSomething() {
	fmt.Println(something)
}
