package main

import (
	"fmt"
	"sort"
)

type Person struct {
	name string
	age  int
}

func main() {
	persons := []Person{
		{"Alejandro", 23},
		{"Eliza", 20},
		{"Alexia", 20},
		{"Chritian", 23},
		{"BKM", 23},
	}

	areSorted := sort.SliceIsSorted(persons, func(i, j int) bool {
		if persons[i].age == persons[j].age {
			return persons[i].name < persons[j].name
		}
		return persons[i].age < persons[j].age
	})

	fmt.Println("are", persons, "sorted? -->", areSorted)

	sort.Slice(persons, func(i, j int) bool {
		if persons[i].age == persons[j].age {
			return persons[i].name < persons[j].name
		}
		return persons[i].age < persons[j].age
	})

	// Esto fue copy paste
	areSorted = sort.SliceIsSorted(persons, func(i, j int) bool {
		if persons[i].age == persons[j].age {
			return persons[i].name < persons[j].name
		}
		return persons[i].age < persons[j].age
	})

	fmt.Println("are", persons, "sorted? -->", areSorted)

}
