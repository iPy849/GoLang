package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

/*
Experiencias:
* Tener cuidado con los canales
* El range + canal se quedaba pegado porque no lo cerraba
 */

var waitingGroup sync.WaitGroup

// Representaciones de datos
type Task struct {
	id int
	data int
}

type TaskResult struct {
	Task
	result int
}

// Canales
var tareas = make(chan Task, runtime.GOMAXPROCS(0))
var resultados = make(chan TaskResult, runtime.GOMAXPROCS(0))

// Funciones
func CreateTask(i int){
	tareas <- Task{id: i, data: i}
}

func RunWorker(){
	for task := range tareas {
		time.Sleep(time.Second)
		resultados <- TaskResult{Task: task, result: task.data * task.data}
	}
	waitingGroup.Done()
	fmt.Println("Termina worker")

}

func ReadResults(){
	for resultado := range resultados {
		fmt.Printf(
			"id: %v\tdata: %v\tresult: %v\n",
			resultado.id,
			resultado.data,
			resultado.result,
			)
	}
}


func main() {
	workers := 1234
	tasksNumber := 10000

	go func() {
		for i := 0; i < tasksNumber; i++ {
			CreateTask(i)
		}
		close(tareas)
	}()

	go ReadResults()


	for i := 0; i < workers; i++ {
		waitingGroup.Add(1)
		go RunWorker()
		fmt.Println("Inicia worker")
	}

	waitingGroup.Wait()
	close(resultados)
	fmt.Println("Terminado")
}