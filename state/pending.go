package main

import (
	"fmt"
	"github.com/pkg/errors"
)

type pending struct {
	product *product
}

func (p *pending) add() error {
	fmt.Println("pending")
	p.product.setStatus(p.product.accepted)
	return nil
}

func (p *pending) assing() error {
	return errors.New("assign error in pending")
}

func (p *pending) done() error {
	return errors.New("done error in pending")
}
