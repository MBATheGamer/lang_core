package parser_test

import (
	"testing"

	"github.com/MBATheGamer/lang_core/lexer"
	"github.com/MBATheGamer/lang_core/parser"
)

func TestSecondLetStatements(t *testing.T) {
	const input = `
let x 5;
let = 10;
let 838383;
	`

	var lexer = lexer.New(input)
	var parser = parser.New(lexer)

	var program = parser.ParseProgram()
	checkParserErrors(t, parser)

	if program == nil {
		t.Fatalf(
			"ParseProgram() returned nil",
		)
	}

	if len(program.Statements) != 3 {
		t.Fatalf(
			"program.Statements does not contain 3 statements. got=%d",
			len(program.Statements),
		)
	}

	var tests = []TestIdentifier{
		{"x"},
		{"y"},
		{"foobar"},
	}

	for i, test := range tests {
		var statement = program.Statements[i]

		if !testLetStatement(t, statement, test.expectedIdentifier) {
			return
		}
	}
}
