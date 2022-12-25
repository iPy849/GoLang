package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

type Record struct {
	Name       string
	Surname    string
	Number     string
	LastAccess string
}

// Registro en memoria de la información
var myData = []Record{}

// Lee archivo con formato csv
func readCSVFile(filepath string) ([][]string, error) {
	// Comprueba que exista el archivo
	_, err := os.Stat(filepath)
	if err != nil {
		return nil, err
	}

	// Abre un acceso al archivo y lo cierra al final de la función
	f, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	// Lee todas las líneas y devuelve un slice bidimensional de strings
	// CSV file read all at once
	// lines data type is [][]string
	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		return [][]string{}, err
	}

	return lines, nil
}

func saveCSVFile(filepath string) error {
	// Crea archivo o lo reescribe y abre acceso a este que cierra al final de la función
	csvfile, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer csvfile.Close()

	// Abre un acceso de escritura al archivo en cuestión
	csvwriter := csv.NewWriter(csvfile)
	// Changing the default field delimiter to tab
	csvwriter.Comma = '\t'
	for _, row := range myData {
		temp := []string{row.Name, row.Surname, row.Number, row.LastAccess}
		_ = csvwriter.Write(temp)
	}
	// Vacía el buffer de escritura
	csvwriter.Flush()
	// Devuelve error nil
	return nil
}

func main() {
	// Requiere un argumento de archivo de entrada y otra de archivo de salida
	if len(os.Args) != 3 {
		fmt.Println("csvData input output!")
		return
	}
	// Comprobación de argumentos
	input := os.Args[1]
	output := os.Args[2]

	// Lee el archivo con formato csv
	lines, err := readCSVFile(input)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Agrega de datos
	// CSV data is read in columns - each line is a slice
	for _, line := range lines {
		temp := Record{
			Name:       line[0],
			Surname:    line[1],
			Number:     line[2],
			LastAccess: line[3],
		}
		myData = append(myData, temp)
		fmt.Println(temp)
	}

	myData = append(myData, Record{
		"Alejandro",
		"Ortega",
		"2101112223",
		"1600665563"},
	)

	// Escribe archivos al disco
	err = saveCSVFile(output)
	if err != nil {
		fmt.Println(err)
		return
	}
}
