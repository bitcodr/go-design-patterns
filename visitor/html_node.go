package main

type htmlNode interface {
	execute(visitor)
}
