package parser

import "github.com/MBATheGamer/lang_core/token"

func (parse *Parser) currentTokenIs(t token.TokenType) bool {
	return parse.currentToken.Type == t
}
