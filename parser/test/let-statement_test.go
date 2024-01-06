package parser_test

import (
	"testing"

	"github.com/MBATheGamer/lang_core/ast"
	"github.com/MBATheGamer/lang_core/lexer"
	"github.com/MBATheGamer/lang_core/parser"
)

func TestLetStatements(t *testing.T) {
	const input = `
let x = 5;
let y = 10;

let foobar = 838383;
	`

	var lex = lexer.New(input)
	var parse = parser.New(lex)

	var program = parse.ParseProgram()

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

func testLetStatement(t *testing.T, statement ast.Statement, name string) bool {
	if statement.TokenLiteral() != "let" {
		t.Errorf(
			"statement.TokenLiteral not 'let'. got=%q",
			statement.TokenLiteral(),
		)

		return false
	}

	var letStatement, ok = statement.(*ast.LetStatement)

	if !ok {
		t.Errorf(
			"statement not *ast.LetStatement. got=%T",
			statement,
		)

		return false
	}

	if letStatement.Name.Value != name {
		t.Errorf(
			"letStatement.Name.Value not '%s'. got=%s",
			name,
			letStatement.Name.Value,
		)

		return false
	}

	if letStatement.Name.TokenLiteral() != name {
		t.Errorf(
			"statement.Name not '%s'. got=%s",
			name,
			letStatement.Name,
		)

		return false
	}

	return true
}
