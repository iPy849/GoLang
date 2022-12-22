package main

import "fmt"

type aStruct struct {
	a int
	b string
}

// TODO: EVITAR PASAR SLICES POR REFERENCIA
// NOTE: EVITAR PASAR SLICES POR REFERENCIA

// Fijate que un apuntador se define usando un * antes del tipo de valor
func myF(a *int, b int) {
	// Valor de apuntador a
	fmt.Println("pointer a:", a)
	// Esto se llama desreferenciar y se trata de extraer el valor al que hace referencia el apuntador
	fmt.Println("value a:", *a)
	// Valor de b
	fmt.Println("value b:", b)
	// Valor de la (dirección de memoria)/puntero de b
	fmt.Println("pointer b:", &b)
}

func myChangePointerValue(a *int) {
	*a++
}

func main() {
	// Valor en memoria
	i := 32
	fmt.Println("value i:", i)
	// Dirección de memoria del valor
	fmt.Println("pointer i:", &i)

	// Extraer la dirección de memoria
	p := &i

	// Punteros y funciones, podemos decir que cuando pasamos un puntero como argumento
	//		podemos decir que lo pasamos por referencia
	myF(p, i)

	// Se demuestra que puedes acceder y modificar el valor de una dirección de memoria
	// 	a través de un puntero
	fmt.Println("value i before changing through a pointer:", i)
	myChangePointerValue(p)
	fmt.Println("value i after changing through a pointer:", i)

	// Al inicializar punteros a tipos de datos, estos se apuntan a ningún lado y EL VALOR CERO DE UN
	//		APUNTADOR ES nil por lo que lo manda a apuntar a nil
	var s *aStruct
	if s == nil {
		fmt.Println("s value is nil:", s)
		// Ocupar un espacio de memoria para la estructura
		s = new(aStruct)
	}

	if s != nil {
		fmt.Println("s value is not nil:", *s, "with pointer", s)
		s.b = "Hola mundo feo y cruel"
		s.a = 2
		fmt.Println("s is now:", *s, "with pointer", s)
	}

}
