package parser_test

import (
	"testing"

	"github.com/MBATheGamer/lang_core/ast"
	"github.com/MBATheGamer/lang_core/lexer"
	"github.com/MBATheGamer/lang_core/parser"
)

func TestParsingIndexExpressions(t *testing.T) {
	var input = "myArray[1 + 1]"

	var lexer = lexer.New(input)
	var parser = parser.New(lexer)
	var program = parser.ParseProgram()
	checkParserErrors(t, parser)

	var statement, _ = program.Statements[0].(*ast.ExpressionStatement)
	var indexExpression, ok = statement.Expression.(*ast.IndexExpression)

	if !ok {
		t.Fatalf(
			"expression not *ast.IndexExpression. got=%T",
			statement.Expression,
		)
	}

	if !testIdentifier(t, indexExpression.Left, "myArray") {
		return
	}

	if !testInfixExpression(t, indexExpression.Index, 1, "+", 1) {
		return
	}
}
