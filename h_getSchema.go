package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/jackc/pgx"
)

func main() {
	arguments := os.Args
	if len(arguments) != 6 {
		fmt.Println("Please provide: hostname port username password db")
		return
	}

	host := arguments[1]
	p := arguments[2]
	user := arguments[3]
	pass := arguments[4]
	database := arguments[5]

	// Port number SHOULD BE an integer
	port, err := strconv.Atoi(p)
	if err != nil {
		fmt.Println("Not a valid port number:", err)
		return
	}

	// open PostgreSQL database with pgx
	conn, err := pgx.Connect(pgx.ConnConfig{
		Host:     host,
		Port:     uint16(port),
		Database: database,
		User:     user,
		Password: pass,
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close()

	// Get all databases
	rows, err := conn.Query(`SELECT "datname" FROM "pg_database"
	WHERE datistemplate = false`)
	if err != nil {
		fmt.Println("Query", err)
		return
	}

	for rows.Next() {
		var name string
		err = rows.Scan(&name)
		defer rows.Close()
		if err != nil {
			fmt.Println("Scan", err)
			return
		}
		fmt.Println("*", name)
	}

	// Get all tables from __current__ database
	query := `SELECT table_name FROM information_schema.tables WHERE 
		table_schema = 'public' ORDER BY table_name`
	rows, err = conn.Query(query)
	if err != nil {
		fmt.Println("Query", err)
		return
	}

	// This is how you process the rows that are returned from SELECT
	for rows.Next() {
		var name string
		err = rows.Scan(&name)
		if err != nil {
			fmt.Println("Scan", err)
			return
		}
		fmt.Println("+T", name)
	}
	defer rows.Close()
}
