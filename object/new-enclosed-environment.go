package object

func NewEnclosedEnivronment(outer *Enivronment) *Enivronment {
	var enivronment = NewEnivronment()
	enivronment.outer = outer
	return enivronment
}
