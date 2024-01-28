package parser_test

import (
	"testing"

	"github.com/MBATheGamer/lang_core/ast"
	"github.com/MBATheGamer/lang_core/lexer"
	"github.com/MBATheGamer/lang_core/parser"
)

func TestIntegerLiteralExpression(t *testing.T) {
	const input = "5;"

	var lexer = lexer.New(input)
	var parser = parser.New(lexer)
	var program = parser.ParseProgram()
	checkParserErrors(t, parser)

	if len(program.Statements) != 1 {
		t.Fatalf(
			"program has not enough statements. got=%d",
			len(program.Statements),
		)
	}

	var statement, _ = program.Statements[0].(*ast.ExpressionStatement)

	var literal, _ = statement.Expression.(*ast.IntegerLiteral)

	if literal.Value != 5 {
		t.Errorf(
			"literal.Value not %d. got=%d",
			5,
			literal.Value,
		)
	}

	if literal.TokenLiteral() != "5" {
		t.Errorf(
			"literal.TokenLiteral not %s. got=%s",
			"5",
			literal.TokenLiteral(),
		)
	}
}
