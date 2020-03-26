package main

import (
	"strconv"
	"fmt"
)

type email struct {
	email string
}

func (e *email) send(key int) {
	fmt.Println("key: " + strconv.Itoa(key) + " sent by email to " + e.email)
}