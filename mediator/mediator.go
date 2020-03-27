package main

type mediator interface {
	communicate(name string)
}

type concreteMediator struct {
	article
	person
}

func newMediator() *concreteMediator {
	mediator := new(concreteMediator)
	mediator.article.setMediator(mediator)
	mediator.person.setMediator(mediator)
	return mediator
}

func (m *concreteMediator) communicate(name string) {
	if name == "nice" {
		m.article.setName("good " + name)
	} else {
		m.person.setName("mike " + name)
	}
}
