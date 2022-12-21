package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	/*
		La regla general para parsear strings a time es que siempre el layout/formato
		va a estar directamente relacionado con los datos:
		Monday 02 January 2006 15:04 MST
		No se puede comparar con más nada para aplicar formato, así lo decidieron cuando
		se escribió Go
	*/

	start := time.Now()

	if len(os.Args) != 2 {
		fmt.Println("usage: dates \"parse_string\"")
		return
	}

	dateString := os.Args[1]

	//Is this date only
	d, err := time.Parse("02 January 2006", dateString)
	if err == nil {
		fmt.Println("Full:", d)
		fmt.Println("Time:", d.Day(), d.Month(), d.Year())
	}

	// Is this a date + time?
	d, err = time.Parse("02 January 2006 15:04", dateString)
	if err == nil {
		fmt.Println("Full:", d)
		fmt.Println("Date:", d.Day(), d.Month(), d.Year())
		fmt.Println("Time:", d.Hour(), d.Minute())
	}

	// Is it time only?
	d, err = time.Parse("15:04", dateString)
	if err == nil {
		fmt.Println("Full:", d)
		fmt.Println("Time:", d.Hour(), d.Minute())
	}

	t := time.Now().Unix()
	fmt.Println("Epoch time:", t)
	// Convert Epoch time to time.Time
	d = time.Unix(t, 0)
	fmt.Println("Date:", d.Day(), d.Month(), d.Year())
	fmt.Printf("Time: %d:%d\n", d.Hour(), d.Minute())
	duration := time.Since(start)
	fmt.Println("Execution time:", duration)

}
