package main


func main() {

	logModel := &log{new(verify)}

	personModel := new(person)
	personModel.email = "a@domain.com"

	logModel.execute(personModel)
}
