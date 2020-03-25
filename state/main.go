package main

import (
	"fmt"
)

//like polymorphism in go
//Use the State design pattern when the object can be in many different states.
//Depending upon current request the object needs to change its current state

func main() {
	initProduct := newProduct("car")

	if err := initProduct.add(); err != nil {
		fmt.Println(err)
	}

	if err := initProduct.assing(); err != nil {
		fmt.Println(err)
	}

	if err := initProduct.done(); err != nil {
		fmt.Println(err)
	}

}
