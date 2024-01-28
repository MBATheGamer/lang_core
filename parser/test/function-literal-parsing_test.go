package parser_test

import (
	"testing"

	"github.com/MBATheGamer/lang_core/ast"
	"github.com/MBATheGamer/lang_core/lexer"
	"github.com/MBATheGamer/lang_core/parser"
)

func TestFunctionLiteralParsing(t *testing.T) {
	const input = "fn (x, y) { x + y; }"

	var lexer = lexer.New(input)
	var parser = parser.New(lexer)
	var program = parser.ParseProgram()
	checkParserErrors(t, parser)

	if len(program.Statements) != 1 {
		t.Fatalf(
			"program.Body does not contain %d statements. got=%d\n",
			1,
			len(program.Statements),
		)
	}

	var statement, _ = program.Statements[0].(*ast.ExpressionStatement)

	var function, _ = statement.Expression.(*ast.FunctionLiteral)

	if len(function.Parameters) != 2 {
		t.Fatalf(
			"function literal parameters wrong. want 2, got=%d\n",
			len(function.Parameters),
		)
	}

	testLiteralExpression(t, function.Parameters[0], "x")
	testLiteralExpression(t, function.Parameters[1], "y")

	if len(function.Body.Statements) != 1 {
		t.Fatalf(
			"function.Body.Statements has not 1 statements. got=%d\n",
			len(function.Body.Statements),
		)
	}

	var bodyStatement, _ = function.Body.Statements[0].(*ast.ExpressionStatement)

	testInfixExpression(t, bodyStatement.Expression, "x", "+", "y")
}
