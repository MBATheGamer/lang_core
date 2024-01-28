package parser_test

import (
	"testing"

	"github.com/MBATheGamer/lang_core/ast"
	"github.com/MBATheGamer/lang_core/lexer"
	"github.com/MBATheGamer/lang_core/parser"
)

func TestBooleanExpression(t *testing.T) {
	var tests = []TestBoolean{
		{"true", true},
		{"false", false},
	}

	for _, test := range tests {
		var lexer = lexer.New(test.input)
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

		var boolean, _ = statement.Expression.(*ast.Boolean)

		if boolean.Value != test.expected {
			t.Errorf(
				"boolean.Value not %t. got=%t",
				test.expected,
				boolean.Value,
			)
		}
	}
}
