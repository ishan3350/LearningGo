package main

import (
	"fmt"
)

type person struct{
	firstname string
	lastname string

	// slice in struct
	favicecreams []string
}

func main(){

	p1 := person{
		firstname: "John",
		lastname: "Test1",
		favicecreams :[]string{"fav1","fav2"},
	}

	p2 := person{
		firstname: "Jack",
		lastname: "Test2",
		favicecreams: []string{"fav3","fav4"},
	}

	fmt.Println(p1.firstname)
	fmt.Println(p1.lastname)

	for favicecream :=  range p1.favicecreams{
		fmt.Println(p1.favicecreams[favicecream])
	}

	fmt.Println(p2.firstname)
	fmt.Println(p2.lastname)

	for favicecream :=  range p2.favicecreams{
		fmt.Println(p2.favicecreams[favicecream])
	}
}
