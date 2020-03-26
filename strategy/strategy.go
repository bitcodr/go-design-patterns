package main

type strategy interface {
	add(name string)
	get() string
}
