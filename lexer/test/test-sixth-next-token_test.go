package lexer_test

import (
	"testing"

	"github.com/MBATheGamer/lang_core/lexer"
	"github.com/MBATheGamer/lang_core/token"
)

func TestSixthNextToken(t *testing.T) {
	const input = `
"foobar"
"foo bar"
	`

	var tests = []TestToken{
		{
			expectedType:    token.STRING,
			expectedLiteral: "foobar",
		},
		{
			expectedType:    token.STRING,
			expectedLiteral: "foo bar",
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
