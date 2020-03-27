package main

import (
	"fmt"
)

type customerCommand struct {
	customerService
}

func (c *customerCommand) execute() {
	fmt.Println("run customer command")
	c.customerService.assingToCustomer()
}
