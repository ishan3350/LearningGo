package main

import (
	"encoding/json"
	"fmt"
)

type User struct{
	First string
	Last string
	Age int
}
func main(){

	u1 := User{
		First: "test",
		Last: "test2",
		Age: 30,
	}

	u2 := User{
		First: "test3",
		Last: "test4",
		Age: 35,
	}

	users := []User{u1,u2}

	data, err := json.Marshal(users)

	if(err == nil){
		fmt.Println(string(data))
	}else{
		fmt.Println(err)
	}
}