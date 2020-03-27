package main

import (
	"fmt"
)

type person struct {
	mediator mediator
	name     string
}

func (p *person) setMediator(m mediator) {
	fmt.Println("set person mediator")
	p.mediator = m
}

func (p *person) setName(name string) {
	fmt.Println("set person name")
	p.name = name
	p.mediator.communicate(name)
}
