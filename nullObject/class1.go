package main

import (
	"fmt"
)

type class1 struct {
	name string
}

func (c1 *class1) numberOfStudents() int {
	fmt.Println(c1.getName() + " has 10 student")
	return 10
}

func (c1 *class1) getName() string {
	return c1.name
}
