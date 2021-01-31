package main

import(
	"fmt"
)

func main(){

	checksport()

}

func checksport(){
	favSport := "Badminton"

	switch favSport{
		case "Badminton":
			fmt.Println("Favourite sport is Badminton")
		case "Cricket":
			fmt.Println("Favourite sport is Cricket")
		default:
			fmt.Println("No conditions are matching")

	}
}