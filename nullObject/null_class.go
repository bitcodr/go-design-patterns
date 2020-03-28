package main

import (
	"fmt"
)

type nullClass struct {
	name string
}

func (n *nullClass) numberOfStudents() int {
	fmt.Println("null object has 0 student")
	return 0
}

func (n *nullClass) getName() string {
	fmt.Println("null object")
	return "null object"
}
