package evaluator

import "github.com/MBATheGamer/lang_core/object"

func evalMinusPrefixOperatorExpression(right object.Object) object.Object {
	if right.Type() != object.INTEGER_OBJ {
		return NULL
	}

	var value = right.(*object.Integer).Value

	return &object.Integer{Value: -value}
}
