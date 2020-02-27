package simple

import (
	"fmt"
	"sort"
)

type (

	//Person struct
	Person struct {
		Name string
		Age  int
	}

	//PersonAscByName person list sort by name
	PersonAscByName []Person
)

//SortMain placeholder to run sort experiments
func SortMain() {
	sortStructPersonByAge()
	sortPersonByName()
}

//Len of persons
func (a PersonAscByName) Len() int {
	return len(a)
}

//Swap of persons
func (a PersonAscByName) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a PersonAscByName) Less(i, j int) bool {
	return a[i].Name <= a[j].Name
}

//SortPersonByName simple way
func sortPersonByName() {
	persons := []Person{
		{"Bob", 12},
		{"Cam", 13},
		{"Alpha", 15},
	}
	sort.Sort(PersonAscByName(persons))
	fmt.Println(persons)
}

//SortStructPersonByAge shortcut
func sortStructPersonByAge() {
	persons := []Person{
		{"Bob", 12},
		{"Cam", 13},
		{"Alpha", 15},
	}
	sort.Slice(persons, func(i, j int) bool {
		return persons[i].Age <= persons[j].Age
	})
	fmt.Println(persons)
}
