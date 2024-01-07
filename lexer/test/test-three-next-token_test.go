package lexer_test

import (
	"testing"

	"github.com/MBATheGamer/lang_core/lexer"
	"github.com/MBATheGamer/lang_core/token"
)

func TestThreeNextToken(t *testing.T) {
	const input = `
let result = add(five, ten);
!-/*5;
5 < 10 > 5;
	`

	var tests = []TestToken{
		{
			expectedType:    token.LET,
			expectedLiteral: "let",
		},
		{
			expectedType:    token.IDENT,
			expectedLiteral: "result",
		},
		{
			expectedType:    token.ASSIGN,
			expectedLiteral: "=",
		},
		{
			expectedType:    token.IDENT,
			expectedLiteral: "add",
		},
		{
			expectedType:    token.LPAREN,
			expectedLiteral: "(",
		},
		{
			expectedType:    token.IDENT,
			expectedLiteral: "five",
		},
		{
			expectedType:    token.COMMA,
			expectedLiteral: ",",
		},
		{
			expectedType:    token.IDENT,
			expectedLiteral: "ten",
		},
		{
			expectedType:    token.RPAREN,
			expectedLiteral: ")",
		},
		{
			expectedType:    token.SEMICOLON,
			expectedLiteral: ";",
		},

		{
			expectedType:    token.BANG,
			expectedLiteral: "!",
		},
		{
			expectedType:    token.MINUS,
			expectedLiteral: "-",
		},
		{
			expectedType:    token.SLASH,
			expectedLiteral: "/",
		},
		{
			expectedType:    token.ASTERISK,
			expectedLiteral: "*",
		},
		{
			expectedType:    token.INT,
			expectedLiteral: "5",
		},
		{
			expectedType:    token.SEMICOLON,
			expectedLiteral: ";",
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
			expectedType:    token.GT,
			expectedLiteral: ">",
		},
		{
			expectedType:    token.INT,
			expectedLiteral: "5",
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
