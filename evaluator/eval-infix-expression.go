package evaluator

import "github.com/MBATheGamer/lang_core/object"

func evalInfixExpression(operator string, left, right object.Object) object.Object {
	switch {
	case left.Type() == object.INTEGER_OBJ && right.Type() == object.INTEGER_OBJ:
		return evalIntegerInfixExpression(operator, left, right)

	case operator == "==":
		return nativeBoolToBooleanObject(left == right)

	case operator == "!=":
		return nativeBoolToBooleanObject(left != right)

	case left.Type() != right.Type():
		return newError(
			"type mismatch: %s %s %s",
			left.Type(),
			operator,
			right.Type(),
		)

	default:
		return newError(
			"unknown operator: %s %s %s",
			left.Type(),
			operator,
			right.Type(),
		)
	}
}
