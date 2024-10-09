package evaluator

import (
	"github.com/MBATheGamer/lang_core/ast"
	"github.com/MBATheGamer/lang_core/object"
)

func evalIfExpression(ifExpression *ast.IfExpression, environment *object.Environment) object.Object {
	var condition = Eval(ifExpression.Condition, environment)

	if isError(condition) {
		return condition
	}

	if isTruthy(condition) {
		return Eval(ifExpression.Consequence, environment)
	} else if ifExpression.Alternative != nil {
		return Eval(ifExpression.Alternative, environment)
	} else {
		return NULL
	}
}
