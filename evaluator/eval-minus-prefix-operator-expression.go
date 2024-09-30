package evaluator

import "github.com/MBATheGamer/lang_core/object"

func evalMinusPrefixOperatorExpression(right object.Object) object.Object {
	if right.Type() != object.INTEGER_OBJ {
		return newError(
			"unknown operator: -%s",
			right.Type(),
		)
	}

	var value = right.(*object.Integer).Value

	return &object.Integer{Value: -value}
}
