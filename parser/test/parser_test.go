package parser_test

import (
	"fmt"
	"testing"

	"github.com/MBATheGamer/lang_core/ast"
	"github.com/MBATheGamer/lang_core/parser"
)

type TestIdentifier struct {
	expectedIdentifier string
}

type TestPrefix struct {
	input        string
	operator     string
	integerValue int64
}

func checkParserError(t *testing.T, parser *parser.Parser) {
	var errors = parser.Errors()

	if len(errors) == 0 {
		return
	}

	t.Errorf(
		"parser has %d errors.",
		len(errors),
	)

	for _, message := range errors {
		t.Errorf(
			"parser error: %q",
			message,
		)
	}

	t.FailNow()
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

func testIntegerLiteral(t *testing.T, il ast.Expression, value int64) bool {
	var integer, ok = il.(*ast.IntegerLiteral)

	if !ok {
		t.Errorf(
			"il not *ast.IntegerLiteral. got=%T",
			il,
		)

		return false
	}

	if integer.Value != value {
		t.Errorf(
			"integer.Value not %d. got=%d",
			value,
			integer.Value,
		)

		return false
	}

	if integer.TokenLiteral() != fmt.Sprintf("%d", value) {
		t.Errorf(
			"integer.TokenLiteral not %d. got=%s",
			value,
			integer.TokenLiteral(),
		)

		return false
	}

	return true
}
