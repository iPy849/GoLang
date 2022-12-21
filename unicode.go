package main

import (
	"fmt"
	"unicode"
)

func main() {
	// Se definen puntos unicode en hex y algunos caracteres ascii intermedios "ab"
	const sL = "\x99\x00ab\x50\x00\x23\x50\x29\x9c"
	// Por cada elemento del slice/array
	for i := 0; i < len(sL); i++ {
		// Si el elemento es imprimible
		if unicode.IsPrint(rune(sL[i])) {
			// Lo imprime como caracter
			fmt.Printf("%c\n", sL[i])
		} else {
			fmt.Println("Not Printable")
		}
	}
}
