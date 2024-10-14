package evaluator

import "github.com/MBATheGamer/lang_core/object"

func applyFunction(fn object.Object, arguments []object.Object) object.Object {
	var function, ok = fn.(*object.Function)

	if !ok {
		return newError("not a function: %s", fn.Type())
	}

	var extendedEnvironment = extendFunctionEnvironment(function, arguments)
	var evaluated = Eval(function.Body, extendedEnvironment)

	return unwrapReturnValue(evaluated)
}
