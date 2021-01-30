package main

import (
	"fmt"

	"github.com/Pallinder/go-randomdata"

	"rsc.io/sampler"
)

func main() {
	fmt.Println(randomdata.SillyName())

	fmt.Println(sampler.Hello())
}
