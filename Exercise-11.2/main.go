package main

import (
	"encoding/json"
	"fmt"
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

	err,bs := toJSON(p1)

	if err == nil{
		fmt.Println(string(bs))

	}else{
		fmt.Println(err)
	}

}

// toJSON needs to return an error also
func toJSON(a interface{}) (error,[]byte) {
	bs, err := json.Marshal(a)

	if err ==nil{
		return nil,bs
	}else{
		// creating custom error using errorf
		return fmt.Errorf("There was ab error in toJSON %v",err),[]byte{}
	}
}
