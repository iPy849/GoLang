package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please, provide an argument")
		return
	}

	file := arguments[1]
	path := os.Getenv("PATH")
	pathSplit := filepath.SplitList(path)
	for _, directory := range pathSplit {

		// Ahora se busca unir el agumento a lo que se busca
		// Esto hace lo mismo que os.path.join en python
		fullPath := filepath.Join(directory, file)

		// Does it exist?
		fileInfo, err := os.Stat(fullPath)
		if err == nil {
			mode := fileInfo.Mode()
			// Is it a regular file?
			if mode.IsRegular() {
				/*
					Al parecer los permisos de usuario no solo se representand con rwx sino
					que también con binarios.El equivalente es:
					r = 0100 => 4
					w = 0010 => 2
					x = 0001 => 1

					Estos permisos pueden combinarse en el word (medio byte) de manera que
					rwx => 0111 => 7
					r-x => 0101 => 5
					rw- => 0110 => 6
					y así se van combinando.

					También los permisos tienen 3 alcances: "owner", "group", "other" en ese mismo orden
					y se agrupan de manera consecutiva, ejemplo:
					$ chmod 644 dirname
				*/
				// Is it executable?
				// Trabaja con los bits
				if mode&0111 != 0 {
					fmt.Println(fullPath)
					return
				}
			}
		}
	}
	fmt.Printf("Cannot find \"%s\" executable\n", file)
}
