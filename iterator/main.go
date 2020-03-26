package main

import (
	"fmt"
)

func main() {

	list := getList()

	for list.hasNext() {
		fmt.Println(list.getItem())
		list.next()
	}
}
