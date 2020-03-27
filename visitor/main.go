package main

func main() {

	anchorModel := new(anchorNode)
	headingModel := new(headingNode)

	highlightModel := new(highlight)
	plainTextModel := new(plainText)
	colorfulModel := new(colorful)

	anchorModel.execute(highlightModel)
	anchorModel.execute(plainTextModel)
	anchorModel.execute(colorfulModel)

	headingModel.execute(highlightModel)
	headingModel.execute(plainTextModel)
	headingModel.execute(colorfulModel)

}
