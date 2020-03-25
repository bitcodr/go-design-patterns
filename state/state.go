package main

type state interface {
	add() error
	assing() error
	done() error
}
