package main

import "fmt"

type admin struct {
	id string
}

func (a *admin) update(itemName string) {
    fmt.Println("admin with id "+ a.id+ " updated by item name "+ itemName)
}

func (a *admin) getID() string {
	return a.id
}
