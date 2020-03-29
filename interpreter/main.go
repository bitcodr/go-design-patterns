package main

import (
	"fmt"
)

func main() {

	isOrderTableSelected := selectTable().interpret("order")

	fmt.Println(isOrderTableSelected)

}
