package main

type observer interface {
	update(itemName string)
	getID() string
}
