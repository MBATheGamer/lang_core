package lexer_test

import (
	"testing"

	"github.com/MBATheGamer/lang_core/lexer"
	"github.com/MBATheGamer/lang_core/token"
)

func TestForthNextToken(t *testing.T) {
	const input = `
if (5 < 10) {
	return true;
} else {
	return false;
}
	`

	var tests = []TestToken{
		{
			expectedType:    token.IF,
			expectedLiteral: "if",
		},
		{
			expectedType:    token.LPAREN,
			expectedLiteral: "(",
		},
		{
			expectedType:    token.INT,
			expectedLiteral: "5",
		},
		{
			expectedType:    token.LT,
			expectedLiteral: "<",
		},
		{
			expectedType:    token.INT,
			expectedLiteral: "10",
		},
		{
			expectedType:    token.RPAREN,
			expectedLiteral: ")",
		},
		{
			expectedType:    token.LBRACE,
			expectedLiteral: "{",
		},

		{
			expectedType:    token.RETURN,
			expectedLiteral: "return",
		},
		{
			expectedType:    token.TRUE,
			expectedLiteral: "true",
		},
		{
			expectedType:    token.SEMICOLON,
			expectedLiteral: ";",
		},

		{
			expectedType:    token.RBRACE,
			expectedLiteral: "}",
		}, {
			expectedType:    token.ELSE,
			expectedLiteral: "else",
		},
		{
			expectedType:    token.LBRACE,
			expectedLiteral: "{",
		},

		{
			expectedType:    token.RETURN,
			expectedLiteral: "return",
		},
		{
			expectedType:    token.FALSE,
			expectedLiteral: "false",
		},
		{
			expectedType:    token.SEMICOLON,
			expectedLiteral: ";",
		},

		{
			expectedType:    token.RBRACE,
			expectedLiteral: "}",
		},

		{
			expectedType:    token.EOF,
			expectedLiteral: "",
		},
	}

	var lex = lexer.New(input)

	for i, test := range tests {
		testToken(t, test, lex, i)
	}
}
