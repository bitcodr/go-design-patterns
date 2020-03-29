package main

type orExpression struct {
	exp1 expression
	exp2 expression
}

func (o *orExpression) interpret(context string) bool {
	return o.exp1.interpret(context) || o.exp2.interpret(context)
}
