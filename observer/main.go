package main

func main() {
	item := newItem("new shoes added")

	customer1 := new(customer)
	customer1.id = "123"
	item.register(customer1)

	customer2 := new(customer)
	customer2.id = "333"
	item.register(customer2)

	item.deregister(customer2)

	admin := new(admin)
	admin.id = "345"
	item.register(admin)

	item.updateAvailable()
}
