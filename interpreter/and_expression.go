package main

type andExpression struct {
	exp1 expression
	exp2 expression
}

func (a *andExpression) interpret(context string) bool {
	return a.exp1.interpret(context) && a.exp2.interpret(context)
}
