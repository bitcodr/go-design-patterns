package main

type product struct {
	pending   state
	accepted  state
	delivered state

	currentState state

	name string
}

func newProduct(name string) *product {
	productModel := new(product)
	productModel.name = name

	pendingModel := &pending{
		product: productModel,
	}
	acceptedModel := &accepted{
		product: productModel,
	}
	deliveredModel := &delivered{
		product: productModel,
	}

	productModel.setStatus(pendingModel)
	productModel.pending = pendingModel
	productModel.accepted = acceptedModel
	productModel.delivered = deliveredModel

	return productModel
}

func (p *product) add() error {
	return p.currentState.add()
}

func (p *product) assing() error {
	return p.currentState.assing()
}

func (p *product) done() error {
	return p.currentState.done()
}

func (p *product) setStatus(s state) {
	p.currentState = s
}
