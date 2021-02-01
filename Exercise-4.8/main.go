package main

import(
	"fmt"
)
// storing slice as value in map

func main(){

	m := map[string][]string{
		"John" :[]string{"hobby1","hobby2","hobby3"},
		"Bill": []string{"hobby4","hobby5","hobby6"},
		"Jack": []string{"hobby7","hobby8","hobby9"},
	}

	for key,val := range m{
		fmt.Println(key)
		for _,hobbies := range val{
			fmt.Println(hobbies)
		}
	}
}