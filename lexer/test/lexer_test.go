package lexer_test

import (
	"testing"

	"github.com/MBATheGamer/lang_core/lexer"
	"github.com/MBATheGamer/lang_core/token"
)

type TestToken struct {
	expectedType    token.TokenType
	expectedLiteral string
}

func testToken(t *testing.T, test TestToken, lex *lexer.Lexer, index int) {
	var token = lex.NextToken()

	if token.Type != test.expectedType {
		t.Fatalf(
			"tests[%d] - token type wrong. expected=%q, got=%q",
			index,
			test.expectedType,
			token.Type,
		)
	}

	if token.Literal != test.expectedLiteral {
		t.Fatalf(
			"tests[%d] - literal wrong. expected=%q, got=%q",
			index,
			test.expectedLiteral,
			token.Literal,
		)
	}
}
