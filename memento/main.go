package main

import "fmt"

func main() {

	stateModel := new(state)
	historyModel := new(history)

	//add state 1
	stateModel.content = "s1"
	s1 := stateModel.create()
	fmt.Println(s1.content)
	historyModel.push(s1)
	fmt.Println(historyModel.states)

	//add state 2
	stateModel.content = "s2"
	s2 := stateModel.create()
	fmt.Println(s2.content)
	historyModel.push(s2)
	fmt.Println(historyModel.states)

	//add state 3
	stateModel.content = "s3"
	s3 := stateModel.create()
	fmt.Println(s3.content)
	historyModel.push(s3)
	fmt.Println(historyModel.states)

	//rollback state 3 to state 2
	prevState2 := stateModel.restore(historyModel.pop())
	fmt.Println(prevState2.content)

	//rollback state 2 to state 1
	stateOfPrev1 := stateModel.restore(historyModel.pop())
	fmt.Println(stateOfPrev1.content)

}
