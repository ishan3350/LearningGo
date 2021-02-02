package main

import(
	"fmt"
)

type vehicle struct{
	doors int
	color string
}

type truck struct{
	vehicle
	isfourwheel bool
}

type sedan struct{
	vehicle
	isluxury bool
}
func main(){

	// struct with in struct
	t1 := truck{
		vehicle: vehicle{
			doors:2,
			color: "red",
		},
		isfourwheel: false,
	}

	// struct inside struct
	s1 := sedan{

		vehicle: vehicle{
			doors: 4,
			color: "Silver",
			

			
		},

		isluxury: true,
	}

	fmt.Println(t1)
	fmt.Println(s1)
	fmt.Println("Truck details:")
	fmt.Println("Vehicle color:\t",t1.vehicle.color)
	fmt.Println("Vehicle color:\t",t1.vehicle.doors)
	fmt.Println("Four Wheel? :\t", t1.isfourwheel)

	fmt.Println("Sedan details:")
	fmt.Println("Vehicle color:\t",s1.vehicle.color)
	fmt.Println("Vehicle color:\t",s1.vehicle.doors)
	fmt.Println("Four Wheel? :\t", s1.isluxury)

}