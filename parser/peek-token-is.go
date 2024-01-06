package parser

import "github.com/MBATheGamer/lang_core/token"

func (parse *Parser) peekTokenIs(t token.TokenType) bool {
	return parse.peekToken.Type == t
}
