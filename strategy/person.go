package main

type person struct {
	strategy strategy
}

func initPerson(s strategy) *person {
	return &person{
		strategy: s,
	}
}

func (p *person) add(name string) {
	p.strategy.add(name)
}

func (p *person) get() string {
	return p.strategy.get()
}
