package evaluator

import "github.com/MBATheGamer/lang_core/object"

func evalIndexExpression(left, index object.Object) object.Object {
	switch {
	case left.Type() == object.ARRAY_OBJ && index.Type() == object.INTEGER_OBJ:
		return evalArrayIndexExpression(left, index)

	case left.Type() == object.HASH_OBJ:
		return evalHashIndexExpression(left, index)

	default:
		return newError(
			"index operator not supported: %s",
			left.Type(),
		)
	}
}
