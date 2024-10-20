package evaluator

import "github.com/MBATheGamer/lang_core/object"

func applyFunction(fn object.Object, arguments []object.Object) object.Object {
	switch fn := fn.(type) {
	case *object.Function:
		var extendedEnvironment = extendFunctionEnvironment(fn, arguments)
		var evaluated = Eval(fn.Body, extendedEnvironment)
		return unwrapReturnValue(evaluated)

	case *object.Builtin:
		return fn.Fn(arguments...)

	default:
		return newError("not a function: %s", fn.Type())
	}
}
