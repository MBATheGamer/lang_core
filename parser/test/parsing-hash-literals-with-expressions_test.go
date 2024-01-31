package parser_test

import (
	"testing"

	"github.com/MBATheGamer/lang_core/ast"
	"github.com/MBATheGamer/lang_core/lexer"
	"github.com/MBATheGamer/lang_core/parser"
)

func TestParsingHashLiteralsWithExpressions(t *testing.T) {
	var input = `{"one": 0 + 1, "two": 10 - 8, "three": 15 / 5}`

	var lexer = lexer.New(input)
	var parser = parser.New(lexer)
	var program = parser.ParseProgram()
	checkParserErrors(t, parser)

	var statement = program.Statements[0].(*ast.ExpressionStatement)
	var hash, _ = statement.Expression.(*ast.HashLiteral)

	if len(hash.Pairs) != 3 {
		t.Errorf(
			"hash.Pairs has wrong length. got=%d",
			len(hash.Pairs),
		)
	}

	var tests = map[string]func(expression ast.Expression){
		"one": func(expression ast.Expression) {
			testInfixExpression(t, expression, 0, "+", 1)
		},
		"two": func(expression ast.Expression) {
			testInfixExpression(t, expression, 10, "-", 8)
		},
		"three": func(expression ast.Expression) {
			testInfixExpression(t, expression, 15, "/", 5)
		},
	}

	for key, value := range hash.Pairs {
		var literal, _ = key.(*ast.StringLiteral)

		var testFunc, _ = tests[literal.String()]

		testFunc(value)
	}
}
