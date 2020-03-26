package main

import (
	"fmt"
)

type customer struct {
	id string
}

func (a *customer) update(itemName string) {
	fmt.Println("customer with id " + a.id + " updated by item name " + itemName)
}

func (a *customer) getID() string {
	return a.id
}
