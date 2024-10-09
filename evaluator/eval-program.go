package evaluator

import (
	"github.com/MBATheGamer/lang_core/ast"
	"github.com/MBATheGamer/lang_core/object"
)

func evalProgram(program *ast.Program, environment *object.Environment) object.Object {
	var result object.Object

	for _, statement := range program.Statements {
		result = Eval(statement, environment)

		switch result := result.(type) {
		case *object.ReturnValue:
			return result.Value

		case *object.Error:
			return result
		}
	}

	return result
}
