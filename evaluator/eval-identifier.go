package evaluator

import (
	"github.com/MBATheGamer/lang_core/ast"
	"github.com/MBATheGamer/lang_core/object"
)

func evalIdentifier(
	node *ast.Identifier,
	environment *object.Environment,
) object.Object {
	if value, ok := environment.Get(node.Value); ok {
		return value
	}

	if builtins, ok := builtins[node.Value]; ok {
		return builtins
	}

	return newError(
		"identifier not found: %s",
		node.Value,
	)
}
