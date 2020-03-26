package main

import "strings"

type lower struct {
	name string
}

func (u *lower) add(name string) {
	u.name = strings.ToLower(name)
}

func (u *lower) get() string {
	return u.name
}
