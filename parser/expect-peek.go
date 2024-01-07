package parser

import "github.com/MBATheGamer/lang_core/token"

func (parser *Parser) expectPeek(tok token.TokenType) bool {
	if parser.peekTokenIs(tok) {
		parser.nextToken()
		return true
	}

	parser.peekError(tok)
	return false
}
