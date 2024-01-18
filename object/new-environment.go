package object

func NewEnivronment() *Enivronment {
	var store = make(map[string]Object)
	return &Enivronment{
		store: store,
		outer: nil,
	}
}
