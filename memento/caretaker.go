package main

type history struct {
	states []*preState
}

func (h *history) push(prevState *preState) {
	h.states = append(h.states, prevState)
}

func (h *history) pop() *preState {
	lastIndex := len(h.states) - 1
	h.states = h.states[:lastIndex]
	preState := new(preState)
	preState.content = h.states[lastIndex-1].content
	return preState
}
