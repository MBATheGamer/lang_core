package object

type Environment struct {
	store map[string]Object
	outer *Environment
}

func (e *Environment) Get(name string) (Object, bool) {
	var object, ok = e.store[name]

	if !ok && e.outer != nil {
		object, ok = e.outer.Get(name)
	}

	return object, ok
}

func (e *Environment) Set(name string, value Object) Object {
	e.store[name] = value
	return value
}
