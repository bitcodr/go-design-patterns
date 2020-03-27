package main

import "fmt"

type log struct {
	next register
}

func (l *log) execute(p *person) {
	fmt.Println("execute log")
	l.next.execute(p)
}