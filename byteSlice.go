package main

import "fmt"

func main() {
	/*
		Aqui lo interesante es que en go no hay un tipo char pero en su lugar
		podemos interpretar un string como un slice de bytes al ser un medida
		estándar entre sistemas digitales, se puede usar para recibir información
		de archivos, strings, etc.

		Al no existir el tipo char lo que se hace es que cada caracter está representado
		por un byte que representa un símbolo ascii o tambien varios bytes pueden
		representar un punto unicode (runa)
	*/
	// Byte slice
	b := make([]byte, 12)
	fmt.Println("Byte slice:", b)
	b = []byte("Byte slice €")
	fmt.Println("Byte slice:", b)

	// Print byte slice contents as text
	fmt.Printf("Byte slice as text: %s\n", b)
	fmt.Println("Byte slice as text:", string(b))

	// Length of b
	fmt.Println("Length of b:", len(b))
}
