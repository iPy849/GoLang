package main

import "errors"

/*
NOTE: My custom implementation of custom errors
*/

import (
	"fmt"
)

func myError(err bool) (int, error) {
	if err == true {
		// Así se levanta un error propio, usa la standar lib "errors"
		return 0, errors.New("This is a new custom error")
	}
	return 1, nil
}

func main() {
	v, e := myError(true)
	if e != nil {
		fmt.Errorf("Ocurrió un error: %s", e)
	}
	fmt.Println(v)
}
