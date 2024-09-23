package evaluator

import (
	"github.com/MBATheGamer/lang_core/ast"
	"github.com/MBATheGamer/lang_core/object"
)

func evalIfExpression(ifExpression *ast.IfExpression) object.Object {
	var condition = Eval(ifExpression.Condition)

	if isError(condition) {
		return condition
	}

	if isTruthy(condition) {
		return Eval(ifExpression.Consequence)
	} else if ifExpression.Alternative != nil {
		return Eval(ifExpression.Alternative)
	} else {
		return NULL
	}
}
