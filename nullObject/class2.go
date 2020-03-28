package main

import (
	"fmt"
)

type class2 struct {
	name string
}

func (c2 *class2) numberOfStudents() int {
	fmt.Println(c2.getName() + " has 12 student")
	return 12
}

func (c2 *class2) getName() string {
	return c2.name
}
