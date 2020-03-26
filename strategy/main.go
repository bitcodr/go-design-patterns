package main

import (
	"fmt"
)

func main() {
	upperModel := new(upper)
	person := initPerson(upperModel)
	person.add("ali")
	fmt.Println(person.get())

	lowerModel := new(lower)
	person = initPerson(lowerModel)
	person.add("REZA")
	fmt.Println(person.get())
}
