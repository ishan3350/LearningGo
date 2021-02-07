package main

import (
	"encoding/json"
	"fmt"
)
type person struct {
	First string `json:"First"`
	Last  string `json:"Last"`
	Age   int    `json:"Age"`
}

func main(){

	data := `[{"First":"Test","Last":"Test2","Age":30},{"First":"test3","Last":"test4","Age":31}]`
	jsonData := []byte(data)
	fmt.Printf("%T",data)
	fmt.Printf("%T",jsonData)
	fmt.Println("Hello World!")

	// creating slice of person type
	people := []person{}

	err := json.Unmarshal(jsonData, &people)

	if err == nil{
		fmt.Println(people)
		// it's basically slice of person type

		fmt.Println(people[0].First)
	}else{
		fmt.Println(err)
	}

}