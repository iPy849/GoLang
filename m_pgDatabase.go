package main

import (
	"database/sql"
	"fmt"

	//NOTE: Se agrega un alias "_" acá para que el compilador no te bote la importación
	//		ya que hay casos donde no usas un librería directamente, sino que una dependencia
	//		activa para otra dependencia
	_ "github.com/lib/pq"
)

var DB_PORT int = 5432
var DB_HOST string = "mel.db.elephantsql.com"
var DB_DATABASE string = "yepqpfwq"
var DB_USER string = "yepqpfwq"
var DB_PASSWORD string = "4g9SXTuWrRCChCoDr4YiOF34NdgDIue5"

func main() {
	//NOTE: La url de conexión que las bd suelen tener no funciona en este caso y requieren
	//		este formato
	connString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		DB_HOST, DB_PORT, DB_USER, DB_PASSWORD, DB_DATABASE)

	db, err := sql.Open("postgres", connString)
	if err != nil {
		fmt.Println("Open()", err)
		panic(err)
	}
	defer db.Close()

	// NOTE: Esto es un query, que si sale bien devuelve un cursor que almacena en rows
	rows, err := db.Query(`SELECT "datname" FROM "pg_database"
	WHERE datistemplate = false`)
	if err != nil {
		fmt.Println("Scan", err)
		panic(err)
	}

	//NOTE: El cursor devuelto carga los datos bajo demanda para no trearse n cantidad de
	//		información simultáneament, para eso está la función Next()
	//TODO: La función Scan() va a recibir tantos punteros como valores pedidos y a va a tratar
	//		de copiarlos
	for rows.Next() {
		var name string
		err = rows.Scan(&name)
		if err != nil {
			fmt.Println("Scan", err)
			return
		}
		fmt.Println("*", name)
	}
	defer rows.Close()
}
