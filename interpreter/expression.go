package main

type expression interface {
	interpret(context string) bool
}

func selectTable() expression {
	selectExpression := new(terminalExpression)
	selectExpression.data = "select"

	tableExpression := new(terminalExpression)
	tableExpression.data = "table order"

	return &andExpression{
		exp1: selectExpression,
		exp2: tableExpression,
	}
}
