package main

import (
	"fmt"
	"os"
	"regexp"
)

func matchInt(s string) bool {
	t := []byte(s)
	re := regexp.MustCompile(`^[-+]?\d+$`)
	return re.Match(t)
}

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Usage: <utility> string.")
		return
	}

	r := [2]int{0, 0}
	for _, v := range arguments[1:] {
		if matchInt(v) == true {
			r[0]++
		} else {
			r[1]++
		}
	}
	fmt.Printf("True: %d False: %d \n", r[0], r[1])
}
