package lexer_test

import (
	"testing"

	"github.com/MBATheGamer/lang_core/lexer"
	"github.com/MBATheGamer/lang_core/token"
)

func TestSeventhNextToken(t *testing.T) {
	const input = `
[1, 2];
	`

	var tests = []TestToken{
		{
			expectedType:    token.LBRACKET,
			expectedLiteral: "[",
		},
		{
			expectedType:    token.INT,
			expectedLiteral: "1",
		},
		{
			expectedType:    token.COMMA,
			expectedLiteral: ",",
		},
		{
			expectedType:    token.INT,
			expectedLiteral: "2",
		},
		{
			expectedType:    token.RBRACKET,
			expectedLiteral: "]",
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
