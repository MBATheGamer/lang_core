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

	var statement, ok = program.Statements[0].(*ast.ExpressionStatement)

	if !ok {
		t.Fatalf(
			"program.Statements[0] is not ast.ExpressionStatement. got=%T",
			program.Statements[0],
		)
	}

	var literal *ast.IntegerLiteral
	literal, ok = statement.Expression.(*ast.IntegerLiteral)

	if !ok {
		t.Fatalf(
			"expression not *ast.IntegerLiteral. got=%T",
			statement.Expression,
		)
	}

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
