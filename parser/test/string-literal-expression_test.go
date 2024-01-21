package parser_test

import (
	"testing"

	"github.com/MBATheGamer/lang_core/ast"
	"github.com/MBATheGamer/lang_core/lexer"
	"github.com/MBATheGamer/lang_core/parser"
)

func TestStringLiteralExpression(t *testing.T) {
	var input = `"Hello, World!"`

	var lexer = lexer.New(input)
	var parser = parser.New(lexer)
	var program = parser.ParseProgram()
	checkParserErrors(t, parser)

	var statement = program.Statements[0].(*ast.ExpressionStatement)
	var literal, ok = statement.Expression.(*ast.StringLiteral)

	if !ok {
		t.Fatalf(
			"expression not *ast.StringLiteral. got=%T",
			statement.Expression,
		)
	}

	if literal.Value != "Hello, World!" {
		t.Errorf(
			"literal.Value not %q. got=%q",
			"Hello, World!",
			literal.Value,
		)
	}
}
