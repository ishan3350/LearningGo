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
		// logging error in a file
		log.SetOutput(f)
		// write error to the log file and exit with status code 1
		log.Fatalln(err)
	} else {
		defer f2.Close()
	}
}
