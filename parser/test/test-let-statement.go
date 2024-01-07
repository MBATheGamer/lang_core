package parser_test

import (
	"testing"

	"github.com/MBATheGamer/lang_core/ast"
)

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
