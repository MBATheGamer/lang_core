package parser

import "github.com/MBATheGamer/lang_core/token"

func (parse *Parser) expectPeek(t token.TokenType) bool {
	if parse.peekTokenIs(t) {
		parse.nextToken()
		return true
	}

	return false
}
