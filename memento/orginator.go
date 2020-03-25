package main

type state struct {
	content string
}

func (s *state) create() *preState {
	preState := new(preState)
	preState.content = s.content
	return preState
}

func (s *state) restore(preState *preState) *state {
	s.content = preState.content
	return s
}
