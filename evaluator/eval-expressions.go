package evaluator

import (
	"github.com/MBATheGamer/lang_core/ast"
	"github.com/MBATheGamer/lang_core/object"
)

func evalExpressions(expressions []ast.Expression, environment *object.Environment) []object.Object {
	var result []object.Object

	for _, expression := range expressions {
		var evaluated = Eval(expression, environment)

		if isError(evaluated) {
			return []object.Object{evaluated}
		}

		result = append(result, evaluated)
	}

	return result
}
