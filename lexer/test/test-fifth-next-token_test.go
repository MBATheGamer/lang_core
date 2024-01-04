package lexer_test

import (
	"testing"

	"github.com/MBATheGamer/lang_core/lexer"
	"github.com/MBATheGamer/lang_core/token"
)

func TestFifthNextToken(t *testing.T) {
	const input = `
10 == 10;
10 != 9;
	`

	var tests = []TestToken{
		{
			expectedType:    token.INT,
			expectedLiteral: "10",
		},
		{
			expectedType:    token.EQ,
			expectedLiteral: "==",
		},
		{
			expectedType:    token.INT,
			expectedLiteral: "10",
		},
		{
			expectedType:    token.SEMICOLON,
			expectedLiteral: ";",
		},

		{
			expectedType:    token.INT,
			expectedLiteral: "10",
		},
		{
			expectedType:    token.NOT_EQ,
			expectedLiteral: "!=",
		},
		{
			expectedType:    token.INT,
			expectedLiteral: "9",
		},
		{
			expectedType:    token.SEMICOLON,
			expectedLiteral: ";",
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
