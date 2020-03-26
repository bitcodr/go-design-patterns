package main

import (
	"strconv"
	"fmt"
)

type sms struct {
    phoneNumber string
}

func (s *sms) send(key int) {
	fmt.Println("key: " + strconv.Itoa(key) + " sent by sms to "+ s.phoneNumber)
}