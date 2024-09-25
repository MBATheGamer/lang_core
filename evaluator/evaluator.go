package evaluator

import (
	"github.com/MBATheGamer/lang_core/ast"
	"github.com/MBATheGamer/lang_core/object"
)

func Eval(node ast.Node) object.Object {
	switch node := node.(type) {
	case *ast.Program:
		return evalStatements(node.Statements)

	case *ast.BlockStatement:
		return evalStatements(node.Statements)

	case *ast.IfExpression:
		return evalIfExpression(node)

	case *ast.ReturnStatement:
		var value = Eval(node.ReturnValue)
		return &object.ReturnValue{Value: value}

	case *ast.ExpressionStatement:
		return Eval(node.Expression)

	case *ast.PrefixExpression:
		var right = Eval(node.Right)
		return evalPrefixExpression(node.Operator, right)

	case *ast.InfixExpression:
		var left = Eval(node.Left)
		var right = Eval(node.Right)
		return evalInfixExpression(node.Operator, left, right)

	case *ast.IntegerLiteral:
		return &object.Integer{Value: node.Value}

	case *ast.Boolean:
		return nativeBoolToBooleanObject(node.Value)
	}

	return nil
}
