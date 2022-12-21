package main

import (
	"fmt"
	"time"
)

func main() {
	// NOTE: Siempre los parsers van a hacer la conversión en base a comparaciones con 02 01 2006 15:04

	// Parsing a time with defined formats
	f1 := "13 03 1999"
	time1, err := time.Parse("02 01 2006", f1)
	if err != nil {
		fmt.Println("Problema en parse:", err)
		return
	}
	fmt.Println(time1)

	// Parsing a time with defined formats
	f2 := "13 Mar 99 00:00 MST"
	time2, err := time.Parse(time.RFC822, f2)
	if err != nil {
		fmt.Println("Problema en parse:", err)
		return
	}
	fmt.Println(time2)

	fmt.Println("Now time Local:", time.Now())
	fmt.Println("Now time UTC:", time.Now().UTC())
	fmt.Println("Now date UTC:", time.Now().Format("02 Jan 2006"))
	fmt.Println("Epoch time:", time.Now().Unix())
}
