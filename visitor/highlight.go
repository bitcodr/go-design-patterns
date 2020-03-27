package main

import (
	"fmt"
)

type highlight struct {
}

func (h *highlight) visitAnchor(a *anchorNode) {
	fmt.Println("highlight anchor")
}

func (h *highlight) visitHeading(a *headingNode) {
	fmt.Println("highlight heading")
}
