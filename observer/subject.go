package main

import "fmt"

type subject interface{
    register(observer)
    deregister(observer)
    notifyObserver()
}


type item struct{
    observers []observer
    name string
    id int
    inStock bool
}


func newItem(name string) *item{
    return &item{
        name: name,
    }
}


func (i *item) updateAvailable() {
    fmt.Println("item in stock")
    i.inStock = true
    i.notifyObserver()
}



func (i *item) register(o observer){
    i.observers = append(i.observers, o)
}


func (i *item) deregister(o observer) {
	for index, value := range i.observers {
		if value.getID() == o.getID() {
            i.observers = append(i.observers[:index], i.observers[index+1:]...)
			break
		}
	}
}

func (i *item) notifyObserver() {
	for _, value := range i.observers {
		value.update(i.name)
	}
}