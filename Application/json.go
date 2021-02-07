package main

// json example
import (
	"encoding/json"
	"fmt"
)

type person struct{
	First string
	Last string
	Age int
}

func main(){

	p1 := person{
		First: "Test",
		Last: "Test2",
		Age: 30,
	}

	p2 := person{
		First: "test3",
		Last:"test4",
		Age: 31,
	}

	// combining two struct objects

	people := []person{p1,p2}
	fmt.Println(people)

	// json encoding or mashaling struct data to json data
	data,err := json.Marshal(people)

	if err == nil{
		fmt.Println(string(data))
	}else{
		fmt.Println(err)
	}
}