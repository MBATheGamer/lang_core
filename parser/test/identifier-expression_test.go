package parser_test

import (
	"testing"

	"github.com/MBATheGamer/lang_core/ast"
	"github.com/MBATheGamer/lang_core/lexer"
	"github.com/MBATheGamer/lang_core/parser"
)

func TestIdentifierExpression(t *testing.T) {
	const input = "foobar;"

	var lexer = lexer.New(input)
	var parser = parser.New(lexer)
	var program = parser.ParseProgram()
	checkParserErrors(t, parser)

	if len(program.Statements) != 1 {
		t.Fatalf(
			"program has not enough statement. got=%d",
			len(program.Statements),
		)
	}

	var statement, _ = program.Statements[0].(*ast.ExpressionStatement)

	var identifier, _ = statement.Expression.(*ast.Identifier)

	if identifier.Value != "foobar" {
		t.Errorf(
			"identifier.Value not %s. got=%s",
			"foobar",
			identifier,
		)
	}

	if identifier.TokenLiteral() != "foobar" {
		t.Errorf(
			"identifier.TokenLiteral not %s. got=%s",
			"foobar",
			identifier.TokenLiteral(),
		)
	}
}
