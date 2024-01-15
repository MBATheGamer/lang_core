package parser_test

import (
	"testing"

	"github.com/MBATheGamer/lang_core/ast"
	"github.com/MBATheGamer/lang_core/lexer"
	"github.com/MBATheGamer/lang_core/parser"
)

func TestThirdLetStatement(t *testing.T) {
	var tests = []TestLetStatement{
		{
			input:              "let x = 5;",
			expectedIdentifier: "x",
			expectedValue:      5,
		},
		{
			input:              "let y = true;",
			expectedIdentifier: "y",
			expectedValue:      true,
		},
		{
			input:              "let foobar = y;",
			expectedIdentifier: "foobar",
			expectedValue:      "y",
		},
	}

	for _, test := range tests {
		var lexer = lexer.New(test.input)
		var parser = parser.New(lexer)
		var program = parser.ParseProgram()
		checkParserErrors(t, parser)

		if len(program.Statements) != 1 {
			t.Fatalf(
				"program.Statements does not contain 1 statement. got=%d",
				len(program.Statements),
			)
		}

		var statement = program.Statements[0]

		if !testLetStatement(t, statement, test.expectedIdentifier) {
			return
		}

		var value = statement.(*ast.LetStatement).Value

		if !testLiteralExpression(t, value, test.expectedValue) {
			return
		}
	}
}
