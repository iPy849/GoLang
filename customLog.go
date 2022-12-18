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
	iLog := log.New(f, "iLog", log.LstdFlags)
	iLog.Println("Hello there!!!")
	iLog.Println("Mastering Go Third Edition!!!")
}
