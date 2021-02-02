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

// type within type example

type occupation struct{
	person
	isselfemployed bool
	jobtitle string
	salary int
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


	// type within type orstruct within struct

	p3 := occupation{
		person: person{
			firstname:"test1",
			lastname:"test2",
			age: 30,
		},
		isselfemployed: false,
		jobtitle: "Information Security Engineer",
		salary: 100000,
	}

	fmt.Println(p3)

	// accessing data

	// get person3's first name

	fmt.Println(p3.firstname)

	// last name
	fmt.Println(p3.lastname)

	// age
	fmt.Println(p3.age)

	// is self employed?
	fmt.Println(p3.isselfemployed)

	// job title
	fmt.Println(p3.jobtitle)

	// salary
	fmt.Println(p3.salary)
}