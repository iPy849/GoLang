package main

import (
	"fmt"
	s "strings" // Una manera de hacer alias
	"unicode"
)

// Otra manera de hacer alias
var f = fmt.Printf

func main() {
	upper := s.ToUpper("Hello there!") // Ámonos a mayúsculas
	f("To Upper: %s\n", upper)
	f("To Lower: %s\n", s.ToLower("Hello THERE")) // Ámonos a minúsculas

	f("%s\n", s.Title("tHis wiLL be A title!")) // NOTE: Obsoleto

	// Compara dos strings sin tener en cuenta mayúsculas y minúsculas
	f("EqualFold: %v\n", s.EqualFold("Mihalis", "MIHAlis"))
	f("EqualFold: %v\n", s.EqualFold("Mihalis", "MIHAli"))

	// NOTE: Esto se explica solo
	f("Prefix: %v\n", s.HasPrefix("Mihalis", "Mi"))
	f("Prefix: %v\n", s.HasPrefix("Mihalis", "mi"))
	f("Suffix: %v\n", s.HasSuffix("Mihalis", "is"))
	f("Suffix: %v\n", s.HasSuffix("Mihalis", "IS"))

	// Devuelve índice de aparición de texto buscado o -1 si no encuentra resultados
	f("Index: %v\n", s.Index("Mihalis", "ha"))
	f("Index: %v\n", s.Index("Mihalis", "Ha"))

	// Numero de apariciones de un substring en un string
	f("Count: %v\n", s.Count("Mihalis", "i"))
	f("Count: %v\n", s.Count("Mihalis", "I"))

	// NOTE: No se va a explicar: esto es 'ab' * 5 en python
	f("Repeat: %s\n", s.Repeat("ab", 5))

	f("TrimSpace: %s\n", s.TrimSpace(" \tThis is a line. \n"))            // Te limpio todos los espacios adicionales
	f("TrimLeft: %s", s.TrimLeft(" \tThis is a\t line. \n", "\n\t "))     // Se entiende que hace
	f("TrimRight: %s\n", s.TrimRight(" \tThis is a\t line. \n", "\n\t ")) // Se entiende que hace

	f("Compare: %v\n", s.Compare("Mihalis", "MIHALIS")) // Compara lexicográficamente dos strings
	f("Compare: %v\n", s.Compare("Mihalis", "Mihalis")) // 1 si a > b; 0 si a = b; -1 si a < b
	f("Compare: %v\n", s.Compare("MIHALIS", "MIHalis")) // La comparación se hace hasta el índice máximo del menor string

	// Se entiend que hace
	t := s.Fields("This is a string!")
	f("Fields: %v\n", len(t))
	// Obviamente el tabulador se le considera como espacio
	t = s.Fields("ThisIs a\tstring!")
	f("Fields: %v\n", len(t))

	// NO voy a explicar esa,
	f("%s\n", s.Split("abcd efg", ""))
	// El replace tiene una sola diferencia, puedes decirle cuantas veces puede efectuar un reemplazo. Si tu n = -1,
	// entonces no ha límites de reemplazoss
	f("%s\n", s.Replace("abcd efg", "", "_", -1))
	f("%s\n", s.Replace("abcd efg", "", "_", 4))
	f("%s\n", s.Replace("abcd efg", "", "_", 2))

	lines := []string{"Line 1", "Line 2", "Line 3"}
	// Se entiende
	f("Join: %s\n", s.Join(lines, "+++"))

	// SplitAfter encuentra un match y hasta ese punto genera un elemento
	f("SplitAfter: %s\n", s.SplitAfter("123++432++", "++"))

	// Esto es como un filtro para de caracteres
	trimFunction := func(c rune) bool {
		return !unicode.IsLetter(c)
	}
	f("TrimFunc: %s\n", s.TrimFunc("123 abc ABC \t .", trimFunction))
}
