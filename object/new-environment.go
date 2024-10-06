package object

func NewEnivronment() *Environment {
	var store = make(map[string]Object)
	return &Environment{
		store: store,
		outer: nil,
	}
}
