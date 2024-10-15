package evaluator

import (
	"github.com/MBATheGamer/lang_core/ast"
	"github.com/MBATheGamer/lang_core/object"
)

func Eval(node ast.Node, environment *object.Environment) object.Object {
	switch node := node.(type) {
	case *ast.Program:
		return evalProgram(node, environment)

	case *ast.LetStatement:
		var value = Eval(node.Value, environment)
		if isError(value) {
			return value
		}
		environment.Set(node.Name.Value, value)

	case *ast.BlockStatement:
		return evalBlockStatement(node, environment)

	case *ast.IfExpression:
		return evalIfExpression(node, environment)

	case *ast.ReturnStatement:
		var value = Eval(node.ReturnValue, environment)
		if isError(value) {
			return value
		}
		return &object.ReturnValue{Value: value}

	case *ast.ExpressionStatement:
		return Eval(node.Expression, environment)

	case *ast.PrefixExpression:
		var right = Eval(node.Right, environment)
		if isError(right) {
			return right
		}
		return evalPrefixExpression(node.Operator, right)

	case *ast.InfixExpression:
		var left = Eval(node.Left, environment)
		if isError(left) {
			return left
		}
		var right = Eval(node.Right, environment)
		if isError(right) {
			return right
		}
		return evalInfixExpression(node.Operator, left, right)

	case *ast.Identifier:
		return evalIdentifier(node, environment)

	case *ast.FunctionLiteral:
		var parameters = node.Parameters
		var body = node.Body
		return &object.Function{
			Parameters:  parameters,
			Enivronment: environment,
			Body:        body,
		}

	case *ast.CallExpression:
		var function = Eval(node.Function, environment)
		if isError(function) {
			return function
		}
		var arguments = evalExpressions(node.Arguments, environment)
		if len(arguments) == 1 && isError(arguments[0]) {
			return arguments[0]
		}
		return applyFunction(function, arguments)

	case *ast.IntegerLiteral:
		return &object.Integer{Value: node.Value}

	case *ast.Boolean:
		return nativeBoolToBooleanObject(node.Value)
	}

	return nil
}
