package object

func NewEnclosedEnivronment(outer *Environment) *Environment {
	var enivronment = NewEnivronment()
	enivronment.outer = outer
	return enivronment
}
