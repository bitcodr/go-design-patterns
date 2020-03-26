package main

import (
	"fmt"
)

func main() {

	key := 1234
	fmt.Println("generate key")

	fmt.Println("save in DB")

	smsModel := new(sms)
	smsModel.phoneNumber = "78987654321"
	sender(smsModel, key)

	emailModel := new(email)
	emailModel.email = "name@domain.com"
	sender(emailModel, key)
}
