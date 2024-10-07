package evaluator

import (
	"github.com/MBATheGamer/lang_core/ast"
	"github.com/MBATheGamer/lang_core/object"
)

func evalIdentifier(node *ast.Identifier, environment *object.Environment) object.Object {
	var value, ok = environment.Get(node.Value)

	if !ok {
		return newError("identifier not found: " + node.Value)
	}

	return value
}
