package main

import (
	"fmt"
	"os"
	"regexp"
)

func matchNameSur(s []string) bool {
	result := true
	for _, v := range s {
		t := []byte(v)
		re := regexp.MustCompile("^[A-Z][a-z]*$")
		result = result && re.Match(t)
	}
	return result
}

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Usage: <utility> string.")
		return
	}
	s := arguments[1:]
	ret := matchNameSur(s)
	fmt.Println(ret)
}
