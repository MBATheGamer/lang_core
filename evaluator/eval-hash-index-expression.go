package evaluator

import "github.com/MBATheGamer/lang_core/object"

func evalHashIndexExpression(
	hash, index object.Object,
) object.Object {
	var hashObject = hash.(*object.Hash)

	var key, ok = index.(object.Hashable)

	if !ok {
		return newError(
			"unusable as hash key: %s",
			index.Type(),
		)
	}

	var pair object.HashPair
	pair, ok = hashObject.Pairs[key.HashKey()]

	if !ok {
		return NULL
	}

	return pair.Value
}
