package main

/*
NOTE: Configuración para este script.

Sencillo, trajimos un paquete del autor del libro donde según el libro
se levantan unos contenedores de docker para una base de datos postgres.

El paquete tiene las operaciones y estructuras para trabajar con el paquete
que se especifica en su repositorio (tiene un .sql con la definición del schema).

Este script es para probar este paquete.

Ahora, por temas de instalación y que meter cualquier software en este equipo
es un problema por los permisos de usuario y hay que llamar a la Karen y no pude
terminar de instalar Docker (hasta el 2023) tuve que improvisar un poco. Solución
fue levantar un servidor externo en un ¿¿¿Database as a Service??? elephantsql.com
y como cliente HeidiSQL. Configuras tus cositas manuales, ajustas las variables de
conexión y voila ya puede correr el script.
*/

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/mactsouk/post05"
)

var MIN = 0
var MAX = 26

func random(min, max int) int {
	return rand.Intn(max-min) + min
}

func getString(length int64) string {
	startChar := "A"
	temp := ""
	var i int64 = 1
	for {
		myRand := random(MIN, MAX)
		newChar := string(startChar[0] + byte(myRand))
		temp = temp + newChar
		if i == length {
			break
		}
		i++
	}
	return temp
}

func main() {
	post05.Hostname = "mel.db.elephantsql.com"
	post05.Port = 5432
	post05.Username = "yepqpfwq"
	post05.Password = "4g9SXTuWrRCChCoDr4YiOF34NdgDIue5"
	post05.Database = "yepqpfwq"

	//NOTE: Llama la función de listar usuarios y los imprime
	data, err := post05.ListUsers()
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, v := range data {
		fmt.Println(v)
	}

	//NOTE: genera información y crea sus datos
	SEED := time.Now().Unix()
	rand.Seed(SEED)
	random_username := getString(5)

	// NOTE: EL resto se entiende relativamente bien, puede echarle un ojo al paquete
	// ya sabes dónde encontrarlos

	t := post05.Userdata{
		Username:    random_username,
		Name:        "Mihalis",
		Surname:     "Tsoukalos",
		Description: "This is me!"}

	id := post05.AddUser(t)
	if id == -1 {
		fmt.Println("There was an error adding user", t.Username)
	}

	err = post05.DeleteUser(id)
	if err != nil {
		fmt.Println(err)
	}

	// Trying to delete it again!
	err = post05.DeleteUser(id)
	if err != nil {
		fmt.Println(err)
	}

	id = post05.AddUser(t)
	if id == -1 {
		fmt.Println("There was an error adding user", t.Username)
	}

	t = post05.Userdata{
		Username:    random_username,
		Name:        "Mihalis",
		Surname:     "Tsoukalos",
		Description: "This might not be me!"}

	err = post05.UpdateUser(t)
	if err != nil {
		fmt.Println(err)
	}
}
