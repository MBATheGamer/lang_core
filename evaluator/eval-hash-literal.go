package evaluator

import (
	"github.com/MBATheGamer/lang_core/ast"
	"github.com/MBATheGamer/lang_core/object"
)

func evalHashLiteral(
	node *ast.HashLiteral,
	environment *object.Environment,
) object.Object {
	var pairs = make(map[object.HashKey]object.HashPair)

	for keyNode, valueNode := range node.Pairs {
		var key = Eval(keyNode, environment)

		if isError(key) {
			return key
		}

		var hashKey, ok = key.(object.Hashable)

		if !ok {
			return newError(
				"unusable as hash key: %s",
				key.Type(),
			)
		}

		var value = Eval(valueNode, environment)

		if isError(value) {
			return value
		}

		var hashed = hashKey.HashKey()
		pairs[hashed] = object.HashPair{
			Key:   key,
			Value: value,
		}
	}

	return &object.Hash{
		Pairs: pairs,
	}
}
