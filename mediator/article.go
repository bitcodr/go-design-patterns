package main

import (
	"fmt"
)

type article struct {
	mediator mediator
	name     string
}

func (a *article) setMediator(m mediator) {
	fmt.Println("set article mediator")
	a.mediator = m
}

func (a *article) setName(name string) {
	fmt.Println("set article name ")
	a.name = name
}

func (a *article) assingPerson(name string) {
	fmt.Println("tell person")
	a.mediator.communicate(name)
}
