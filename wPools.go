package main

import (
	"fmt"
	"os"
	"runtime"
	"strconv"
	"sync"
	"time"
)

/*
	Este script está bien chido, implementa un pool de tareas (la traducción
	es rara: piscina de tareas jajaja).

	La idea aquí es crear una lista de tareas que va a ser trabajada por workers
	y cuando estos terminen toman otra tarea, es una manera de mantener la carga
	de procesamiento controlada.

	Es algo parecido a lo que pasa con los hilos de procesamiento y con los servidores
	web (al menos apache funciona bajo el mismo principio)
 */

// Representa un trabajo
type Client struct {
	id      int
	integer int
}

// Representa un trabajo terminado
type Result struct {
	job    Client
	square int
}

// Inicializa los canales para comunicar gorrutinas
var size = runtime.GOMAXPROCS(0)
var clients = make(chan Client, size)
var data = make(chan Result, size)

func worker(wg *sync.WaitGroup) {
	for c := range clients {
		square := c.integer * c.integer
		output := Result{c, square}
		data <- output
		time.Sleep(time.Second)
	}
	wg.Done()
}

// Crear n cantidad de solicitudes
func create(n int) {
	for i := 0; i < n; i++ {
		c := Client{i, i}
		clients <- c
	}
	close(clients)
}

func main() {
	// region trata el cli

	if len(os.Args) != 3 {
		fmt.Println("Need #jobs and #workers!")
		return
	}

	nJobs, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	nWorkers, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println(err)
		return
	}
	// endregion

	// Crea las solicitudes de manera asíncrona
	go create(nJobs)

	// bandera de terminación
	finished := make(chan interface{})
	// Mientras el canal data no este cerrado, espera cada resultado para imprimirlo en pantalla
	go func() {
		for d := range data {
			fmt.Printf("Client ID: %d\tint: ", d.job.id)
			fmt.Printf("%d\tsquare: %d\n", d.job.integer, d.square)
		}
		finished <- true
	}()

	var wg sync.WaitGroup
	// Lanza un worker con una referencia al waitingGroup para que avisen de terminar
	for i := 0; i < nWorkers; i++ {
		wg.Add(1)
		go worker(&wg)
	}
	wg.Wait()
	close(data)

	fmt.Printf("Finished: %v\n", <-finished)
}
