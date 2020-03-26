package main

type iterator interface {
	hasNext() bool
	next()
	getItem() string
}

func getList() iterator {
	return &stringList{
		list:  []string{"a", "b", "c"},
		count: 0,
	}
}
