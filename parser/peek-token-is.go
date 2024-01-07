package parser

import "github.com/MBATheGamer/lang_core/token"

func (parser *Parser) peekTokenIs(tok token.TokenType) bool {
	return parser.peekToken.Type == tok
}
