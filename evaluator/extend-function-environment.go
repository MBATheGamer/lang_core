package evaluator

import "github.com/MBATheGamer/lang_core/object"

func extendFunctionEnvironment(
	fn *object.Function,
	arguments []object.Object,
) *object.Environment {
	var environment = object.NewEnclosedEnivronment(fn.Enivronment)

	for parameterIndex, parameter := range fn.Parameters {
		environment.Set(parameter.Value, arguments[parameterIndex])
	}

	return environment
}
