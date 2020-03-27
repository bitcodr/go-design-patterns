package main

import (
	"fmt"
)


type shop struct {
	command command
}

func (s *shop) buy() {
	fmt.Println("start shopping")
	s.command.execute()
}