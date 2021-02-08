package main

import (
	"fmt"
	"sort"
)

// custom sort
type Person struct {
	Name string
	Age  int
}
type ByAge []Person

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

func main() {

	p1 := Person{
		Name: "charlie",
		Age:  20,
	}

	p2 := Person{
		Name: "bravo",
		Age:  10,
	}

	p3 := Person{
		Name: "test",
		Age:  15,
	}

	// combining all struct object in one slice

	people := []Person{p1, p2, p3}

	// sorting slice by name
	sort.Slice(people, func(i, j int) bool { return people[i].Name < people[j].Name })

	fmt.Println(people)

	// sorting  data in slice by Age
	sort.Sort(ByAge(people))

	fmt.Println(people)

}
