package main

import (
	"fmt"
	"log"
	"os"
	"path"
)

func main() {
	// Este print es para no perder el rastro del log
	fmt.Println(path.Join(os.TempDir(), "mGo.log"))
	LOGFILE := path.Join(os.TempDir(), "mGo.log")
	// Carga un archivo en los modos de Abrir, Crear o Adjuntar
	f, err := os.OpenFile(LOGFILE, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}

	// defer indica que cuando se acabe la función se ejecute la función que se aplicó el defer
	defer f.Close()

	// Cargas tu nuevo archivo de log y añades información a este
	// En este caso las diferentes banderas que les carguemos producen cambios en los logs

	LstdFlags := log.Ldate | log.Lshortfile       //Fecha y nombre corto
	iLog := log.New(f, "LNum ", LstdFlags)        // Carga la función del logger
	iLog.Println("Mastering Go, 3rd edition!")    // Imprime log
	iLog.SetFlags(log.Lshortfile | log.LstdFlags) // Cambia las banderas
	iLog.Println("Another log entry!")            // Imprime log
}
