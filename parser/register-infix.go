package parser

import "github.com/MBATheGamer/lang_core/token"

func (parser *Parser) registerInfix(tokenType token.TokenType, fn infixParseFn) {
	parser.infixParseFns[tokenType] = fn
}
