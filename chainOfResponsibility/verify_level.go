package main

import (
	"fmt"
)

type verify struct {
	next register
}

func (l *verify) execute(p *person) {
	fmt.Println("execute verify")
}
