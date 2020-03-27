package main

import (
	"fmt"
)

type colorful struct {
}

func (h *colorful) visitAnchor(a *anchorNode) {
	fmt.Println("colorful anchor")
}

func (h *colorful) visitHeading(a *headingNode) {
	fmt.Println("colorful heading")
}
