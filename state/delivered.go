package main

import (
	"fmt"
	"github.com/pkg/errors"
)

type delivered struct {
	product *product
}

func (p *delivered) add() error {
	return errors.New("add error in deliver")

}

func (p *delivered) assing() error {
	return errors.New("assign error in deliver")
}

func (p *delivered) done() error {
	fmt.Println("deliver")
	p.product.setStatus(p.product.delivered)
	return nil
}
