package main

import (
	"chapter5/mpkg"
	"chapter5/mpkg2"
	"fmt"

	"github.com/google/uuid"
)

/*
NOTE: Ahora que lo empiezo a entender no hay tanto lío, de todos modos
dale un ojo a README.md que esta vez está bueno jajaja
*/
func main() {
	id := uuid.New()
	fmt.Println(id)
	mpkg.SayHello()
	mpkg.ReadmeSomething()
	fmt.Println(mpkg.YELL)
	mpkg2.SayHello()
}
