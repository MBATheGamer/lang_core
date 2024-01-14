package parser_test

import (
	"testing"

	"github.com/MBATheGamer/lang_core/ast"
	"github.com/MBATheGamer/lang_core/lexer"
	"github.com/MBATheGamer/lang_core/parser"
)

func TestFunctionParameterParsing(t *testing.T) {
	var tests = []TestFunctionParameters{
		{
			input:    "fn() {};",
			expected: []string{},
		},
		{
			input:    "fn(x) {};",
			expected: []string{"x"},
		},
		{
			input:    "fn(x, y, z) {};",
			expected: []string{"x", "y", "z"},
		},
	}

	for _, test := range tests {
		var lexer = lexer.New(test.input)
		var parser = parser.New(lexer)
		var program = parser.ParseProgram()
		checkParserErrors(t, parser)

		var statement = program.Statements[0].(*ast.ExpressionStatement)
		var function = statement.Expression.(*ast.FunctionLiteral)

		if len(function.Parameters) != len(test.expected) {
			t.Errorf(
				"length parameters wrong. want %d, got=%d\n",
				len(test.expected),
				len(function.Parameters),
			)
		}

		for i, identifier := range test.expected {
			testLiteralExpression(t, function.Parameters[i], identifier)
		}
	}
}
