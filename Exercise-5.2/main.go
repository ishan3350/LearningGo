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

	m := map[string]person{
		p1.lastname: p1,
		p2.lastname: p2,
	}

	for key,val := range m{
		fmt.Printf("\nLast name: %v",key)
		fmt.Printf("\nFirst name: %v", val.firstname)

		for index,data := range val.favicecreams{
			fmt.Printf("\nFav icecream %d:\t%v ",index+1,data)
		}
	}
}
