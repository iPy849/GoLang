package main

import "fmt"

func main() {
	aString := "Hola Amigo!¢"
	fmt.Println("original string:", aString)

	// Se descompone el string a slice de byte
	aByteSlice := []byte(aString)
	fmt.Println("byte slice:", aByteSlice)
	for _, b := range aByteSlice {
		// Se observa representación entera y en string de cada rune
		rune := rune(b)
		fmt.Printf("%v -> %s\n", rune, string(rune))
	}
	fmt.Println()
	// Rune inicialization
	rune := 'µ'
	fmt.Printf("rune -> raw value: %v int value: %d rune value: %c type: %T\n", rune, rune, rune, rune)
}
