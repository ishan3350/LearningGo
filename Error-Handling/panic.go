package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Create("log.txt")

	if err != nil {
		fmt.Println(err)

		// log.Println will also show date and time stamp
		log.Println(err)
	} else {
		defer f.Close()
		fmt.Println("Log file created")
	}

	f2, err := os.Open("testing.txt")

	if err != nil {
		panic(err)
	} else {
		defer f2.Close()
	}

}
