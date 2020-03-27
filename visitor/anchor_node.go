package main

type anchorNode struct {
}

func (a *anchorNode) execute(o visitor) {
	o.visitAnchor(a)
}
