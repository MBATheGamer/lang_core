package parser

import "github.com/MBATheGamer/lang_core/token"

func (parser *Parser) registerPrefix(tokenType token.TokenType, fn prefixParseFn) {
	parser.prefixParseFns[tokenType] = fn
}
