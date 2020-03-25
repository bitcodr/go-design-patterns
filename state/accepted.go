package main

import (
	"fmt"
	"github.com/pkg/errors"
)

type accepted struct {
	product *product
}

func (p *accepted) add() error {
	return errors.New("error add in accept")
}

func (p *accepted) assing() error {
	fmt.Println("accept")
	p.product.setStatus(p.product.delivered)
	return nil
}

func (p *accepted) done() error {
	return errors.New("error done in accept")
}
