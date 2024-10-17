package evaluator

import "github.com/MBATheGamer/lang_core/object"

func evalStringInfixExpression(
	operator string,
	left, right object.Object,
) object.Object {
	if operator != "+" {
		return newError(
			"unknown operator: %s %s %s",
			left.Type(),
			operator,
			right.Type(),
		)
	}

	var leftValue = left.(*object.String).Value
	var rightValue = right.(*object.String).Value

	return &object.String{Value: leftValue + rightValue}
}
