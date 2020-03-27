package main

import (
	"fmt"
)

type plainText struct {
}

func (h *plainText) visitAnchor(a *anchorNode) {
	fmt.Println("plain text anchor")
}

func (h *plainText) visitHeading(a *headingNode) {
	fmt.Println("plain text heading")
}
