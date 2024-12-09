package evaluator

import (
	"github.com/MBATheGamer/lang_core/ast"
	"github.com/MBATheGamer/lang_core/object"
)

func evalBlockStatement(
	block *ast.BlockStatement,
	environment *object.Environment,
) object.Object {
	var result object.Object

	for _, statement := range block.Statements {
		result = Eval(statement, environment)

		if result != nil {
			var resultType = result.Type()

			if resultType == object.RETURN_VALUE_OBJ || resultType == object.ERROR_OBJ {
				return result
			}
		}
	}

	return result
}
