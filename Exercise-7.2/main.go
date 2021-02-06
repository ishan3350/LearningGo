package main

// changing value of struct object using pointer

import (
	"fmt"
)

type person struct {
	first string
	last  string
	age   int
}

func changevalue(p *person) {
	*&p.first = "test3"
	*&p.last = "test4"
	*&p.age = 31
}
func main() {

	a := 10
	fmt.Println(&a)

	p1 := person{
		first: "test",
		last:  "test2",
		age:   30,
	}

	fmt.Println(p1)

	changevalue(&p1)

	fmt.Println(p1)

}
