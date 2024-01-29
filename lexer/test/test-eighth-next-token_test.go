package lexer_test

import (
	"testing"

	"github.com/MBATheGamer/lang_core/lexer"
	"github.com/MBATheGamer/lang_core/token"
)

func TestEighthNextToken(t *testing.T) {
	const input = `
{"foo": "bar"};
	`

	var tests = []TestToken{
		{
			expectedType:    token.LBRACE,
			expectedLiteral: "{",
		},
		{
			expectedType:    token.STRING,
			expectedLiteral: "foo",
		},
		{
			expectedType:    token.COLON,
			expectedLiteral: ":",
		},
		{
			expectedType:    token.STRING,
			expectedLiteral: "bar",
		},
		{
			expectedType:    token.RBRACE,
			expectedLiteral: "}",
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

	var lexer = lexer.New(input)

	for i, test := range tests {
		testToken(t, test, lexer, i)
	}
}
