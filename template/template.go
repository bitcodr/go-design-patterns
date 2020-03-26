package main

type template interface {
	send(key int)
}

func sender(t template, key int) {
	t.send(key)
}
