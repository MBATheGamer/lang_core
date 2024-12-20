package evaluator

import "github.com/MBATheGamer/lang_core/object"

func evalBangOperatorExpression(
	right object.Object,
) object.Object {
	switch right {
	case TRUE:
		return FALSE
	case FALSE:
		return TRUE
	case NULL:
		return TRUE
	default:
		return FALSE
	}
}
