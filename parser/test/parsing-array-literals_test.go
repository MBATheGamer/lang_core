package parser_test

import (
	"testing"

	"github.com/MBATheGamer/lang_core/ast"
	"github.com/MBATheGamer/lang_core/lexer"
	"github.com/MBATheGamer/lang_core/parser"
)

func TestParsingArrayLiterals(t *testing.T) {
	var input = "[1, 2 * 2, 3 + 3]"

	var lexer = lexer.New(input)
	var parser = parser.New(lexer)
	var program = parser.ParseProgram()
	checkParserErrors(t, parser)

	var statement, _ = program.Statements[0].(*ast.ExpressionStatement)
	var array, _ = statement.Expression.(*ast.ArrayLiteral)

	if len(array.Elements) != 3 {
		t.Fatalf(
			"len(array.Elements) not 3. got=%d",
			len(array.Elements),
		)
	}

	testIntegerLiteral(t, array.Elements[0], 1)
	testInfixExpression(t, array.Elements[1], 2, "*", 2)
	testInfixExpression(t, array.Elements[2], 3, "+", 3)
}
