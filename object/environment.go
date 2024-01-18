package object

type Enivronment struct {
	store map[string]Object
	outer *Enivronment
}

func (e *Enivronment) Get(name string) (Object, bool) {
	var object, ok = e.store[name]

	if !ok && e.outer != nil {
		object, ok = e.outer.Get(name)
	}

	return object, ok
}

func (e *Enivronment) Set(name string, value Object) Object {
	e.store[name] = value
	return value
}
