package evaluator

import "github.com/MBATheGamer/lang_core/object"

func evalIntegerInfixExpression(operator string, left, right object.Object) object.Object {
	var leftValue = left.(*object.Integer).Value
	var rightValue = right.(*object.Integer).Value

	switch operator {
	case "+":
		return &object.Integer{Value: leftValue + rightValue}
	case "-":
		return &object.Integer{Value: leftValue - rightValue}
	case "*":
		return &object.Integer{Value: leftValue * rightValue}
	case "/":
		return &object.Integer{Value: leftValue / rightValue}

	case "<":
		return nativeBoolToBooleanObject(leftValue < rightValue)
	case ">":
		return nativeBoolToBooleanObject(leftValue > rightValue)
	case "==":
		return nativeBoolToBooleanObject(leftValue == rightValue)
	case "!=":
		return nativeBoolToBooleanObject(leftValue != rightValue)

	default:
		return newError(
			"unknown operator: %s %s %s",
			left.Type(),
			operator,
			right.Type(),
		)
	}
}
