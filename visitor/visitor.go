package main

type visitor interface {
	visitHeading(*headingNode)
	visitAnchor(*anchorNode)
}
