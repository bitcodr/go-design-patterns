package main

import "strings"

type upper struct {
	name string
}

func (u *upper) add(name string) {
	u.name = strings.ToUpper(name)
}

func (u *upper) get() string {
	return u.name
}
