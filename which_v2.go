package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func searchExecutable(file string) {
	// Se busca la variable de entorna $PATH
	path := os.Getenv("PATH")
	// Se hace un slice usando como separador el separador del sistema
	pathSplit := filepath.SplitList(path)
	// Bandera para saber si se tuvo resultados
	var found bool

	for _, directory := range pathSplit {
		// Ahora se busca unir el nombre del ejecutable buscado con cada directorio del path
		// Esto hace lo mismo que os.path.join en python
		fullPath := filepath.Join(directory, file)

		// Does it exist?
		fileInfo, err := os.Stat(fullPath)
		if err == nil {
			mode := fileInfo.Mode()
			// Is it a regular file?
			if mode.IsRegular() {
				// Is it executable?
				if mode&0111 != 0 {
					fmt.Println(fullPath)
					found = true
				}
			}
		}
	}
	if !found {
		fmt.Printf("Cannot find \"%s\" executable\n", file)
	}
}

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please, provide an argument")
		return
	}
	for _, file := range arguments[1:] {
		// Separador por búsqueda
		fmt.Printf("---- %s ----\n", file)
		// Ejecutar búsqueda
		searchExecutable(file)
	}
}
