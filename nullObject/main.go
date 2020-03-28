package main

func main() {

	schoolModel := new(school)

	class1Model := new(class1)
	class1Model.name = "class1"

	schoolModel.setClass(class1Model)

	class2Model := new(class2)
	class2Model.name = "class2"

	schoolModel.getClass(class2Model).numberOfStudents()

}
