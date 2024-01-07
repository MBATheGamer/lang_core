package parser

import (
	"github.com/MBATheGamer/lang_core/lexer"
	"github.com/MBATheGamer/lang_core/token"
)

type Parser struct {
	lexer        *lexer.Lexer
	currentToken token.Token
	peekToken    token.Token
	errors       []string
}

func New(lex *lexer.Lexer) *Parser {
	var parser = &Parser{
		lexer:  lex,
		errors: []string{},
	}

	parser.nextToken()
	parser.nextToken()

	return parser
}
