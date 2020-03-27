package main

func main() {

	customerServiceModel := NewCustomerService("mike")
	customerCommand := &customerCommand{customerServiceModel}

	shopModel := &shop{
		command: customerCommand,
	}
	shopModel.buy()

}
