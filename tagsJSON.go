package main

import (
	"encoding/json"
	"fmt"
)

/*
NOTE Fíjate en esto:

type NoEmpty struct {
	Name    string `json:"username"`
	Surname string `json:"surname"`
	Year    int    `json:"creationyear,omitempty"`
}

Es una estructura de Go muy corriente como cualquier otra pero esta tiene metadatos,
esos metadatos son los `json: "<name>,[options]"` que trae segudio del tipo.

¿Cuál es la utilidad de esto? Esto permite al marshaling (proceso de convertir estructuras
de go en formato json) recibir información en su procesos. En el caso anterior tenemos que
el campo Name de la estructura será traducido a "username" en el json y de la misma forma
con el campo Surname, pero en el campo Year decimos que queremos mostrarlo como
"creationyear" y que lo queremos omitir si está vacío.
*/

// Ignoring empty fields in JSON
type NoEmpty struct {
	Name    string `json:"username"`
	Surname string `json:"surname"`
	Year    int    `json:"creationyear,omitempty"`
}

// Removing private fields and ignoring empty fields
type Password struct {
	Name    string `json:"username"`
	Surname string `json:"surname,omitempty"`
	Year    int    `json:"creationyear,omitempty"`
	Pass    string `json:"-"` // En este caso le estamos diciendo que no lo muestre
}

func main() {
	noempty := NoEmpty{Name: "Mihalis"}
	password := Password{Name: "Mihalis", Pass: "myPassword"}

	// Ignoring empty fields in JSON
	noEmptyVar, err := json.Marshal(&noempty)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("noEmptyVar decoded with value %s\n", noEmptyVar)
	}

	// Removing private fields
	passwordVar, err := json.Marshal(&password)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("password decoded with value %s\n", passwordVar)
	}
}
