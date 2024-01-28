package parser_test

import (
	"testing"

	"github.com/MBATheGamer/lang_core/ast"
	"github.com/MBATheGamer/lang_core/lexer"
	"github.com/MBATheGamer/lang_core/parser"
)

func TestCallExpressionParsing(t *testing.T) {
	const input = "add(1, 2 * 3, 4 + 5);"

	var lexer = lexer.New(input)
	var parser = parser.New(lexer)
	var program = parser.ParseProgram()
	checkParserErrors(t, parser)

	if len(program.Statements) != 1 {
		t.Fatalf(
			"program.Statements does not contain %d statements. got=%d\n",
			1,
			len(program.Statements),
		)
	}

	var statement, _ = program.Statements[0].(*ast.ExpressionStatement)

	var expression, _ = statement.Expression.(*ast.CallExpression)

	if !testIdentifier(t, expression.Function, "add") {
		return
	}

	if len(expression.Arguments) != 3 {
		t.Fatalf(
			"wrong length of arguments. got=%d",
			len(expression.Arguments),
		)
	}

	testLiteralExpression(t, expression.Arguments[0], 1)
	testInfixExpression(t, expression.Arguments[1], 2, "*", 3)
	testInfixExpression(t, expression.Arguments[2], 4, "+", 5)
}
