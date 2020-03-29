package main

import (
	"strings"
)

type terminalExpression struct {
	data string
}

func (t *terminalExpression) interpret(context string) bool {
	return strings.Contains(t.data, context)
}
