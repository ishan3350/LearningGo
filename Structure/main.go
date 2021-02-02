package main

import (
	"fmt"
)

// struct example

type person struct{
	firstname string
	lastname string
	age int
}
func main(){

	p1 := person{
		firstname: "Jack",
		lastname: "Awesome",
		age: 20,
	}

	fmt.Println(p1)

	//printing first name

	fmt.Println(p1.firstname)

	// printing last name
	fmt.Println(p1.lastname)

	// printing age
	fmt.Println(p1.age)

	p2 := person{
		firstname: "Bill",
		lastname: "Amazing",
		age: 30,
	}

	fmt.Println(p2)

	//printing first name

	fmt.Println(p2.firstname)

	// printing last name
	fmt.Println(p2.lastname)

	// printing age
	fmt.Println(p2.age)


}