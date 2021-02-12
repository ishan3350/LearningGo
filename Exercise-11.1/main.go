package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type person struct {
	First   string
	Last    string
	Sayings []string
}

func main() {
	p1 := person{
		First:   "James",
		Last:    "Bond",
		Sayings: []string{"Shaken, not stirred", "Any last wishes?", "Never say never"},
	}

	bs, err:= json.Marshal(p1)
	fmt.Println(string(bs))

	if err ==nil{
		fmt.Println(bs)
	}else{
		log.Fatalln(err)
	}

}
