package evaluator

import (
	"github.com/MBATheGamer/lang_core/ast"
	"github.com/MBATheGamer/lang_core/object"
)

func evalStatements(statements []ast.Statement, environment *object.Environment) object.Object {
	var result object.Object

	for _, statement := range statements {
		result = Eval(statement, environment)

		if returnValue, ok := result.(*object.ReturnValue); ok {
			return returnValue.Value
		}
	}

	return result
}
