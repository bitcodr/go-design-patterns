package main

type headingNode struct {
}

func (h *headingNode) execute(o visitor) {
	o.visitHeading(h)
}
