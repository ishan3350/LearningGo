package main

import (
	"fmt"
)

func main(){
	simplemapexample()

}
func simplemapexample(){
	// map

	//map is key and value pair and it allow fast and efficient search. Name and age mapping is the example


	/* syntax mapname = map[key-datatype]value-datatype{

		key: value,
		key:value,
		key:value,

	}

	*/
	m := map[string]int{
		"James": 10,
		"John": 15,
		"Bill":35,
	}

	fmt.Println(m)

	// accessing map elements

	// finding age of Bill
	fmt.Println(m["Bill"])

	/* if key is not found in map then it will return default value of the value type so for example if we are trying to find the age of a person that doesn't exist in 
	this map then it will return int default value or zero value 0 */

	// how to check if the key exist in map or not

	val,status := m["person1"]

	// As you can see person1 doesn't exist in map so it will return default value 0 and status boolean false
	fmt.Println(val,status)

	// checking if the key exist or not before printing it
	if data,ok := m["Bill"]; ok{
		fmt.Println(data)
	}

	// adding new key, value element in map

	m["Jack"] = 20

	fmt.Println(m)

	// looping through map using for loop and range

	for keydata, valuedata := range m{
		fmt.Println(keydata,valuedata)
	} 


}