package lexer_test

import (
	"testing"

	"github.com/MBATheGamer/lang_core/lexer"
	"github.com/MBATheGamer/lang_core/token"
)

func TestSecondNextToken(t *testing.T) {
	const input = `
let five = 5;
let ten = 10;

let add = fn(x, y) {
	x + y;
};
	`

	var tests = []TestToken{
		{
			expectedType:    token.LET,
			expectedLiteral: "let",
		},
		{
			expectedType:    token.IDENT,
			expectedLiteral: "five",
		},
		{
			expectedType:    token.ASSIGN,
			expectedLiteral: "=",
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
			expectedType:    token.LET,
			expectedLiteral: "let",
		},
		{
			expectedType:    token.IDENT,
			expectedLiteral: "ten",
		},
		{
			expectedType:    token.ASSIGN,
			expectedLiteral: "=",
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
			expectedType:    token.LET,
			expectedLiteral: "let",
		},
		{
			expectedType:    token.IDENT,
			expectedLiteral: "add",
		},
		{
			expectedType:    token.ASSIGN,
			expectedLiteral: "=",
		},
		{
			expectedType:    token.FUNCTION,
			expectedLiteral: "fn",
		},
		{
			expectedType:    token.LPAREN,
			expectedLiteral: "(",
		},
		{
			expectedType:    token.IDENT,
			expectedLiteral: "x",
		},
		{
			expectedType:    token.COMMA,
			expectedLiteral: ",",
		},
		{
			expectedType:    token.IDENT,
			expectedLiteral: "y",
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
			expectedType:    token.IDENT,
			expectedLiteral: "x",
		},
		{
			expectedType:    token.PLUS,
			expectedLiteral: "+",
		},
		{
			expectedType:    token.IDENT,
			expectedLiteral: "y",
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
