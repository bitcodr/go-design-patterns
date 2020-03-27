package main

import (
	"fmt"
)

type customerService interface {
	assingToCustomer()
}

type customer struct {
	name string
}

func NewCustomerService(name string) customerService {
	return &customer{
		name: name,
	}
}

func (c *customer) assingToCustomer() {
	fmt.Println("shop assigned to customer " + c.name)
}
