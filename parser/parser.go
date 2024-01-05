package parser

import (
	"github.com/MBATheGamer/lang_core/lexer"
	"github.com/MBATheGamer/lang_core/token"
)

type Parser struct {
	lex          *lexer.Lexer
	currentToken token.Token
	peekToken    token.Token
}

func New(lex *lexer.Lexer) *Parser {
	var parse = &Parser{lex: lex}

	parse.nextToken()
	parse.nextToken()

	return parse
}
