package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func check(a, b int) error {
	if a == 0 && b == 0 {
		return errors.New("This is a new custom error")
	}
	return nil
}

func formattedError(a, b int) error {
	if a == 0 && b == 0 {
		return fmt.Errorf("a %d and b %d. User ID: %d", a, b, os.Getuid())
	}
	return nil
}

func main() {

	// Calls check and ends successfully
	err := check(0, 10)
	if err == nil {
		fmt.Println("check() ended normally!")
	} else {
		fmt.Println(err)
	}

	// Calls check and ends with error
	err = check(0, 0)
	if err.Error() == "This is a new custom error" {
		fmt.Println("Custom error detected!")
	}

	// Calls formattedError and ends with error
	err = formattedError(0, 0)
	if err != nil {
		fmt.Println(err)
	}

	// Calls strconv.Atoi and ends with successfully
	i, err := strconv.Atoi("-123")
	if err == nil {
		fmt.Println("Int value is", i)
	}

	// Calls strconv.Atoi and ends with error
	i, err = strconv.Atoi("Y123")
	if err != nil {
		fmt.Println(err)
	}
}
